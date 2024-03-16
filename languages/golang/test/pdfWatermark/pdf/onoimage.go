package pdf

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"lgo/test/unipdf/common"
	"lgo/test/unipdf/contentstream"
	"lgo/test/unipdf/core"
	"lgo/test/unipdf/creator"
	"lgo/test/unipdf/extractor"
	"lgo/test/unipdf/model"
	"os"
)

func TestMain() {
	inputPath := "test.pdf"
	//测试获取文字
	text, err := GetPdfTextContent(inputPath, 1)
	if err != nil {
		return
	}
	fmt.Println(text)

	//测试获取图片
	imageContent, err := GetPdfImageContent(inputPath, 1)
	if err != nil {
		return
	}
	if len(imageContent) > 0 {
		// testWriteFile(imageContent[0])
	}
}

// GetPdfTextContent 获取pdf文字内容,起始页pageNum = 1
func GetPdfTextContent(inputPath string, pageNum int) (string, error) {
	pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
	if err != nil {
		return "", err
	}
	defer f.Close()

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
			fmt.Println(imageName)
			if imageName != "Img1" {

				usedObjectsNames = append(usedObjectsNames, imageName)
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
