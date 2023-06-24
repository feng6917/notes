package main

import (
	"errors"
	"fmt"
	"lgo/unipdf/creator"
	"lgo/unipdf/model"
	"lgo/util"
	"lgo/util/file"

	// "lgo/unipdf/contentstream/draw"
	// "lgo/unipdf/core"
	"flag"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

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

// 读取默认配置
func init() {
	flag.BoolVar(&defaultValue, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-", false, tip)
	flag.StringVar(&inputFilePath, "i", "", "name: 输入文件路径. \r\nvalue: 必传,不能为空.")
	flag.StringVar(&outputFilePath, "o", "", "name: 输出文件路径. \r\nvalue: 必传,不能为空. ")
	flag.StringVar(&watermarkFilePath, "w", "", "name: 水印文件路径. \r\nvalue: 必传,不能为空. ")
	flag.StringVar(&startPage, "s", "", "name: 开始页数. \r\nvalue: 非必传,默认首页. ")
	flag.StringVar(&endPage, "e", "", "name: 结束页数. \r\nvalue: 非必传,默认尾页.")
	flag.StringVar(&proof, "p", "", "name: 是否校对. \r\nvalue: (t/f)非必传,默认不校对.")
	flag.StringVar(&ran, "r", "", "name: 是否随机. \r\nvalue: (t/f)非必传,默认不随机.")
}

func main() {
	flag.Parse()
	var err error
	if strings.TrimSpace(inputFilePath) == "" || strings.TrimSpace(outputFilePath) == "" || strings.TrimSpace(watermarkFilePath) == "" {
		fmt.Println(util.ErrorStr("文件路径不能为空"))
		return
	}

	// 校验文件路径
	if !file.PathExists(inputFilePath) || !file.PathExists(watermarkFilePath) {
		fmt.Println(util.ErrorStr("文件不存在"))
		return
	}

	var startPageInt, endPageInt int
	if startPage != "" {
		startPageInt, err = strconv.Atoi(startPage)
		if err != nil {
			fmt.Println(util.ErrorStr(err.Error()))
			return
		}
	}

	if endPage != "" {
		endPageInt, err = strconv.Atoi(endPage)
		if err != nil {
			fmt.Println(util.ErrorStr(err.Error()))
			return
		}
	}
	var proofStatus, rangeStatus bool
	if proof != "" {
		if proof == "t" {
			proofStatus = true
		} else {
			if proof != "f" {
				fmt.Println(util.ErrorStr("校对值无法解析"))
				return
			}
		}
	}

	if ran != "" {
		if ran == "t" {
			rangeStatus = true
		} else {
			if ran != "f" {
				fmt.Println(util.ErrorStr("随机值无法解析"))
				return
			}
		}
	}

	err = addImageToPdf(inputFilePath, outputFilePath, watermarkFilePath, startPageInt, endPageInt, proofStatus, rangeStatus)
	if err != nil {
		fmt.Printf("Error: addImageToPdf %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputFilePath)
}

func addImageToPdf(inputPath string, outputPath string, watermarkFilePath string, startPage, endPage int, proofStaus, rangeStatus bool) error {
	c := creator.New()

	// Prepare the image.
	watermarkImg, err := c.NewImageFromFile(watermarkFilePath)
	if err != nil {
		fmt.Printf("Error: c.NewImageFromFile %v\n", err)
		return err
	}

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

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		fmt.Printf("Error: pdfReader.GetNumPages %v\n", err)
		return err
	}

	if startPage > numPages {
		return errors.New("开始页数过大")
	}

	endPage -= 1

	if endPage > numPages-1 || endPage <= 0 {
		endPage = numPages - 1
	}

	if startPage > 0 {
		startPage -= 1
	}

	// Load the pages.
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			fmt.Printf("Error: pdfReader.GetPage %v\n", err)
			return err
		}

		// Add the page.
		err = c.AddPage(page)
		if err != nil {
			fmt.Printf("Error: c.AddPage %v\n", err)
			return err
		}

		// If the specified page, or -1, apply the image to the page.
		if i >= startPage && i <= endPage {
			if proofStaus {
				watermarkImg.ScaleToWidth(c.Context().PageWidth)
			}
			w := c.Context().PageHeight - watermarkImg.Width()
			h := c.Context().PageHeight - watermarkImg.Height()
			var x, y float64
			if rangeStatus {
				x, y = getRandFloat(w/2), getRandFloat(h/2)
			}
			// 设置位置
			watermarkImg.SetPos(x, y)
			// 设置清晰度
			watermarkImg.SetOpacity(0.5)
			_ = c.Draw(watermarkImg)
		}
	}

	err = c.WriteToFile(outputPath)
	return err
}

func getRandFloat(v float64) float64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	r := rand.New(rand.NewSource(time.Now().UnixNano() + randSeek))
	res := r.Intn(int(v))
	return float64(res)

}
