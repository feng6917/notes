package pdf

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"lgo/test/unipdf/common"
	"lgo/test/unipdf/contentstream"
	"lgo/test/unipdf/core"
	"lgo/test/unipdf/creator"
	"lgo/test/unipdf/extractor"
	"lgo/test/unipdf/model"
	"os"
	"sync"
)

var imageCache *sync.Map

func TestMain() {
	inputPath := "test.pdf"
	// 获取pdf页码
	pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
	if err != nil {
		return
	}
	defer f.Close()

	//测试获取文字
	text, err := GetPdfTextContent(pdfReader, 2)
	if err != nil {
		return
	}
	if text != "" {
		fmt.Println("==================== text")
		fmt.Println(text)
		// page, err := pdfReader.GetPage(2)
		// if err != nil {
		// 	return
		// }
		// textXobjects(page)
	}

	// //测试获取图片
	// imageContent, err := GetPdfImageContent(inputPath, 1)
	// if err != nil {
	// 	return
	// }
	// if len(imageContent) > 0 {
	// 	testWriteFile(imageContent[0], 1)
	// }
}

func testImage(pdfReader *model.PdfReader) {
	pageCount, err := pdfReader.GetNumPages()
	if err != nil {
		return
	}
	imageCache = &sync.Map{}
	for i := 0; i < pageCount; i++ {
		summaryImageCount(pdfReader, i+1)
	}
	c1 := creator.New()
	// c1.AddPage(page)
	for i := 0; i < pageCount; i++ {
		// fmt.Println(i)
		tmp, err := newImage(pdfReader, i+1)
		if err == nil {
			// fmt.Println(i)
			c1.AddPage(tmp)
		}
	}
	c1.WriteToFile("123.pdf")
}

// GetPdfTextContent 获取pdf文字内容,起始页pageNum = 1
func GetPdfTextContent(pdfReader *model.PdfReader, pageNum int) (string, error) {

	page, err := pdfReader.GetPage(pageNum)
	//导出文本
	extract, err := extractor.New(page)
	if err != nil {
		return "", err
	}
	text, err := extract.ExtractText()
	if err != nil {
		return "", err
	}

	return text, nil
}

// GetPdfImageContent 获取pdf文件图片
func GetPdfImageContent(inputPath string, pageNum int) ([][]byte, error) {
	// 打开pdf文件
	pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	page, err := pdfReader.GetPage(pageNum)
	extract, err := extractor.New(page)
	if err != nil {
		return nil, err
	}
	//获取unipdf图片内容
	args := &extractor.ImageExtractOptions{
		IncludeInlineStencilMasks: false,
	}
	images, err := extract.ExtractPageImages(args)
	if err != nil {
		return nil, err
	}

	//转化为go 图片处理并获取图片byte
	var ImageBytes [][]byte
	for index, img := range images.Images {
		//处理单个图片
		oneImage, err := GetOnePdfImageContent(img)
		if err != nil {
			return nil, err
		}
		ImageBytes = append(ImageBytes, oneImage)
		testWriteFile(oneImage, index)
	}
	cleanUnusedXobjects(page)
	c1 := creator.New()
	c1.AddPage(page)
	c1.WriteToFile("123.pdf")
	return ImageBytes, nil
}

// GetOnePdfImageContent 处理单个pdf图片
func GetOnePdfImageContent(imageBuff extractor.ImageMark) ([]byte, error) {
	//转换为go image
	goImage, err := imageBuff.Image.ToGoImage()
	if err != nil {
		return nil, err
	}
	//图片编码
	buf := new(bytes.Buffer)
	opt := jpeg.Options{Quality: 100}
	err = jpeg.Encode(buf, goImage, &opt)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func testWriteFile(data []byte, index int) {
	//写入文件
	fileWrite, err := os.OpenFile(fmt.Sprintf("./liu_%d.jpg", index), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	fileWrite.Write(data)
}

func cleanUnusedXobjects(page *model.PdfPage) {
	contents, err := page.GetAllContentStreams()
	if err != nil {
		common.Log.Debug("failed to get page content stream")
	}
	parser := contentstream.NewContentStreamParser(contents)
	operations, err := parser.Parse()
	if err != nil {
		common.Log.Debug("failed to parse content stream")
	}
	usedObjectsNames := []string{}
	for _, op := range *operations {
		operand := op.Operand
		// Check for `Do` (Draw XObject) operator.
		if operand == "Do" {
			params := op.Params
			imageName := params[0].String()
			val, valExist := imageCache.Load(imageName)
			if !valExist {
				usedObjectsNames = append(usedObjectsNames, imageName)
			} else {
				valInt := val.(int)
				// TODO: 待处理
				if valInt < 10 {
					usedObjectsNames = append(usedObjectsNames, imageName)
				}
			}
		}
	}

	xObject := page.Resources.XObject
	dict, ok := xObject.(*core.PdfObjectDictionary)
	if ok {
		keys := getKeys(dict)
		for _, k := range keys {
			if exists(k, usedObjectsNames) {
				continue
			}
			name := *core.MakeName(k)
			dict.Remove(name)
		}
	}

}

func getKeys(dict *core.PdfObjectDictionary) []string {
	keys := []string{}
	for _, k := range dict.Keys() {
		keys = append(keys, k.String())
	}
	return keys
}

func exists(element string, elements []string) bool {
	for _, el := range elements {
		if element == el {
			return true
		}
	}
	return false
}

func summaryImageCount(pdfReader *model.PdfReader, pageNum int) error {
	page, err := pdfReader.GetPage(pageNum)
	if err != nil {
		return err
	}
	summaryXobjects(page)
	return nil
}

func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func newImage(pdfReader *model.PdfReader, pageNum int) (*model.PdfPage, error) {
	page, err := pdfReader.GetPage(pageNum)
	if err != nil {
		return nil, err
	}
	cleanUnusedXobjects(page)

	return page, nil
}

func summaryXobjects(page *model.PdfPage) {
	contents, err := page.GetAllContentStreams()
	if err != nil {
		common.Log.Debug("failed to get page content stream")
	}
	parser := contentstream.NewContentStreamParser(contents)
	operations, err := parser.Parse()
	if err != nil {
		common.Log.Debug("failed to parse content stream")
	}
	for _, op := range *operations {
		operand := op.Operand
		// Check for `Do` (Draw XObject) operator.
		if operand == "Do" {
			params := op.Params
			key := params[0].String()
			cacheVal, cacheExist := imageCache.Load(key)
			if cacheExist {
				valInt := cacheVal.(int)
				imageCache.Store(key, valInt+1)
			} else {
				imageCache.Store(key, 1)
			}
		}
	}
}

func textXobjects(page *model.PdfPage) {
	contents, err := page.GetAllContentStreams()
	if err != nil {
		common.Log.Debug("failed to get page content stream")
	}
	parser := contentstream.NewContentStreamParser(contents)
	operations, err := parser.Parse()
	if err != nil {
		common.Log.Debug("failed to parse content stream")
	}
	usedObjectsNames := []string{}
	for _, op := range *operations {
		fmt.Println("params: ", op.Params)
		fmt.Println("Operand: ", op.Operand)
		// operand := op.Operand
		// Check for `Do` (Draw XObject) operator.
		// if operand == "Do" {
		// 	params := op.Params
		// 	imageName := params[0].String()
		// 	fmt.Println("imageName: ",imageName)
		// }
	}

	xObject := page.Resources.XObject
	dict, ok := xObject.(*core.PdfObjectDictionary)
	if ok {
		keys := getKeys(dict)
		for _, k := range keys {
			if exists(k, usedObjectsNames) {
				continue
			}
			name := *core.MakeName(k)
			dict.Remove(name)
		}
	}

}
