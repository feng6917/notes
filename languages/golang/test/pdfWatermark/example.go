package main

import (

	// "lgo/unipdf/contentstream/draw"
	// "lgo/unipdf/core"

	"sync"
)

// TODO: 页面版水印添加工具 待调整

const (
	tip string = "PDF 添加水印操作: \r\napp -i input.pdf -o output.pdf -w watermark.jpg \r\n负杂操作: \napp -i input.pdf -o output.pdf -w watermark.jpg -s 3 -e 5 -r t -p t"
)

var (
	inputFilePath     string // 输入文件路径
	outputFilePath    string // 输出文件路径
	watermarkFilePath string // 水印文件路径
	startPage         string // 开始页数
	endPage           string // 结束页数
	proof             string // 校对
	ran               string // 随机
	defaultValue      bool

	randSeek = int64(1)
	l        sync.Mutex
)

// func main() {
// flag.Parse()
// var err error
// if strings.TrimSpace(inputFilePath) == "" || strings.TrimSpace(outputFilePath) == "" || strings.TrimSpace(watermarkFilePath) == "" {
// 	fmt.Println(util.ErrorStr("文件路径不能为空"))
// 	return
// }

// // 校验文件路径
// if !file.PathExists(inputFilePath) || !file.PathExists(watermarkFilePath) {
// 	fmt.Println(util.ErrorStr("文件不存在"))
// 	return
// }

// var startPageInt, endPageInt int
// if startPage != "" {
// 	startPageInt, err = strconv.Atoi(startPage)
// 	if err != nil {
// 		fmt.Println(util.ErrorStr(err.Error()))
// 		return
// 	}
// }

// if endPage != "" {
// 	endPageInt, err = strconv.Atoi(endPage)
// 	if err != nil {
// 		fmt.Println(util.ErrorStr(err.Error()))
// 		return
// 	}
// }
// var proofStatus, rangeStatus bool
// if proof != "" {
// 	if proof == "t" {
// 		proofStatus = true
// 	} else {
// 		if proof != "f" {
// 			fmt.Println(util.ErrorStr("校对值无法解析"))
// 			return
// 		}
// 	}
// }

// if ran != "" {
// 	if ran == "t" {
// 		rangeStatus = true
// 	} else {
// 		if ran != "f" {
// 			fmt.Println(util.ErrorStr("随机值无法解析"))
// 			return
// 		}
// 	}
// }

// err = addImageToPdf(inputFilePath, outputFilePath, watermarkFilePath, startPageInt, endPageInt, proofStatus, rangeStatus)
// if err != nil {
// 	fmt.Printf("Error: addImageToPdf %v\n", err)
// 	os.Exit(1)
// }

// fmt.Printf("Complete, see output file: %s\n", outputFilePath)
// runAPP()
// }

// func addImageToPdf(inputPath string, outputPath string, watermarkFilePath string, startPage, endPage int, proofStaus, rangeStatus bool) error {
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
// 			if proofStaus {
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

// func getRandFloat(v float64) float64 {
// 	l.Lock()
// 	if randSeek >= 100000000 {
// 		randSeek = 1
// 	}
// 	randSeek++
// 	l.Unlock()
// 	r := rand.New(rand.NewSource(time.Now().UnixNano() + randSeek))
// 	res := r.Intn(int(v))
// 	return float64(res)

// }
