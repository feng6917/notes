package main

import (
	"os"
	"strings"

	"lgo/test/pdfWatermark/pdf"
	"lgo/test/pdfWatermark/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/flopp/go-findfont"
)

func init() {
	//设置中文字体:解决中文乱码问题
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	// pdf.ListImages("test.pdf")
	// runAPP()
	pdf.TestMain()
}

func runAPP() {
	myApp := app.New()

	myApp.Settings().SetTheme(theme.LightTheme())
	w := myApp.NewWindow("PDF&WaterMark")

	// service
	svc := service.Service{}
	svc.Init()
	content := svc.Run(myApp, w)

	// 运行应用
	w.SetContent(content)
	w.Resize(fyne.NewSize(850, 200))
	w.ShowAndRun()
}
