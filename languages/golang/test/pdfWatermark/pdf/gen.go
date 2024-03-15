package pdf

import (
	"bytes"
	"fmt"
	"lgo/test/pdfWatermark/imageutil"
	"lgo/test/unipdf/creator"
	"lgo/test/unipdf/model"
	"os"

	"github.com/duke-git/lancet/v2/random"
)

type Pdf struct {
	creator   *creator.Creator
	image     *creator.Image
	pdfReader *model.PdfReader
}

func (self *Pdf) Init() {
	self.creator = creator.New()
}

func (self *Pdf) InitImage(watermarkFilePath string, sliderThumbnail float64) error {
	var imageBuf []byte
	var err error
	imageBuf, err = os.ReadFile(watermarkFilePath)
	if err != nil {
		return err
	}
	resBuf := bytes.NewBuffer([]byte{})
	
	_, _, err = imageutil.Scale(bytes.NewReader(imageBuf), resBuf, sliderThumbnail, 100)
	if err != nil {
		return err
	}
	image, err := self.creator.NewImageFromData(resBuf.Bytes())

	if err != nil {
		return err
	}
	self.image = image
	return nil
}

func (self *Pdf) InitPdfReader(inputPath string) error {
	// Read the input pdf file.
	f, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error: os.Open %v\n", err)
		return err
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		fmt.Printf("Error: model.NewPdfReader %v\n", err)
		return err
	}

	self.pdfReader = pdfReader
	return nil
}

func (self *Pdf) Do(singleton, rangeStatus bool, sliderClarity float64, progressChan chan float64) error {
	numPages, err := self.pdfReader.GetNumPages()
	if err != nil {
		fmt.Printf("Error: pdfReader.GetNumPages %v\n", err)
		return err
	}

	// Load the pages.
	for i := 0; i < numPages; i++ {
		page, err := self.pdfReader.GetPage(i + 1)
		if err != nil {
			fmt.Printf("Error: pdfReader.GetPage %v\n", err)
			return err
		}

		// Add the page.
		err = self.creator.AddPage(page)
		if err != nil {
			fmt.Printf("Error: c.AddPage %v\n", err)
			return err
		}

		// If the specified page, or -1, apply the image to the page.
		if i >= 0 && i <= numPages {
			if singleton {
				self.writeSingleImage(rangeStatus, sliderClarity)
			} else {
				self.writeAnyImage(rangeStatus, sliderClarity)
			}
		}
		go func() {
			progressChan <- float64(i) / float64(numPages)
		}()
	}
	return nil
}

func (self *Pdf) writeSingleImage(rangeStatus bool, slider float64) {
	w := self.creator.Context().PageHeight - self.image.Width()
	h := self.creator.Context().PageHeight - self.image.Height()
	var x, y float64
	if rangeStatus && w > 0 && h > 0 {
		x, y = random.RandFloat(0, w/2, 2), random.RandFloat(0, h/2, 2)
	}
	// 设置位置
	self.image.SetPos(x, y)
	// 设置清晰度
	self.image.SetOpacity(slider)
	_ = self.creator.Draw(self.image)
}

func (self *Pdf) writeAnyImage(rangeStatus bool, slider float64) {
	hights := []float64{0}
	var tmpH float64
	for {
		tmpH += self.image.Height()
		if tmpH < self.creator.Context().PageHeight {
			hights = append(hights, tmpH)
		} else {
			break
		}
	}
	fs := random.RandFloats(len(hights), 0, self.creator.Context().PageWidth-self.image.Width(), 2)
	for index, i := range hights {
		var x, y float64
		if rangeStatus {
			x, y = fs[index], i
		}
		// 设置位置
		self.image.SetPos(x, y)
		// 设置清晰度
		self.image.SetOpacity(slider)
		_ = self.creator.Draw(self.image)
	}
}

func (self *Pdf) Save(outPath string) error {
	err := self.creator.WriteToFile(outPath)
	return err
}

// func AddImageToPdf(inputPath string, outputPath string, watermarkFilePath string, startPage, endPage int, proofStatus, rangeStatus bool) error {
// 	c := creator.New()

// 	// Prepare the image.
// 	watermarkImg, err := c.NewImageFromFile(watermarkFilePath)
// 	if err != nil {
// 		fmt.Printf("Error: c.NewImageFromFile %v\n", err)
// 		return err
// 	}

// 	// Read the input pdf file.
// 	f, err := os.Open(inputPath)
// 	if err != nil {
// 		fmt.Printf("Error: os.Open %v\n", err)
// 		return err
// 	}
// 	defer f.Close()

// 	pdfReader, err := model.NewPdfReader(f)
// 	if err != nil {
// 		fmt.Printf("Error: model.NewPdfReader %v\n", err)
// 		return err
// 	}

// 	numPages, err := pdfReader.GetNumPages()
// 	if err != nil {
// 		fmt.Printf("Error: pdfReader.GetNumPages %v\n", err)
// 		return err
// 	}

// 	if startPage > numPages {
// 		return errors.New("开始页数过大")
// 	}

// 	endPage -= 1

// 	if endPage > numPages-1 || endPage <= 0 {
// 		endPage = numPages - 1
// 	}

// 	if startPage > 0 {
// 		startPage -= 1
// 	}

// 	// Load the pages.
// 	for i := 0; i < numPages; i++ {
// 		page, err := pdfReader.GetPage(i + 1)
// 		if err != nil {
// 			fmt.Printf("Error: pdfReader.GetPage %v\n", err)
// 			return err
// 		}

// 		// Add the page.
// 		err = c.AddPage(page)
// 		if err != nil {
// 			fmt.Printf("Error: c.AddPage %v\n", err)
// 			return err
// 		}

// 		// If the specified page, or -1, apply the image to the page.
// 		if i >= startPage && i <= endPage {
// 			if proofStatus {
// 				watermarkImg.ScaleToWidth(c.Context().PageWidth)
// 			}
// 			w := c.Context().PageHeight - watermarkImg.Width()
// 			h := c.Context().PageHeight - watermarkImg.Height()
// 			var x, y float64
// 			if rangeStatus {
// 				x, y = getRandFloat(w/2), getRandFloat(h/2)
// 			}
// 			// 设置位置
// 			watermarkImg.SetPos(x, y)
// 			// 设置清晰度
// 			watermarkImg.SetOpacity(0.5)
// 			_ = c.Draw(watermarkImg)
// 		}
// 	}

// 	err = c.WriteToFile(outputPath)
// 	return err
// }

// func compressImageResource(data []byte) []byte {
// 	imgSrc, _, err := image.Decode(bytes.NewReader(data))
// 	if err != nil {
// 		return data
// 	}

// 	newImg := resize.Resize(0, 200, imgSrc, resize.Lanczos3)
// 	buf := bytes.Buffer{}
// 	err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 100})
// 	if err != nil {
// 		return data
// 	}
// 	if buf.Len() > len(data) {
// 		return data
// 	}
// 	return buf.Bytes()
// }

// func png2jpg(pngPath string) ([]byte, error) {
// 	f, err := os.Open(pngPath)
// 	if err != nil {
// 		log.Errorf("os open failed, err: %v", err)
// 		return nil, err
// 	}
// 	defer f.Close()
// 	img, err := png.Decode(f)
// 	if err != nil {
// 		log.Errorf("png decode failed, err: %v", err)
// 		return nil, err
// 	}
// 	var buf bytes.Buffer
// 	// 将图像转换为JPEG格式并写入文件
// 	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
// 	if err != nil {
// 		log.Errorf("jpeg encode failed, err: %v", err)
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }
