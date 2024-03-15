package service

import (
	"lgo/test/pdfWatermark/fileutil"
	"lgo/test/pdfWatermark/progress"
	"lgo/test/pdfWatermark/selectutil"
	"lgo/test/pdfWatermark/slider"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Service struct {
	fileInPdf       fileutil.FileInPdf
	fileInWater     fileutil.FileInWater
	fileOut         fileutil.FileOut
	selectRand      selectutil.SelectRandom
	selectCount     selectutil.SelectSingleton
	sliderClarity   slider.SliderClarity
	sliderThumbnail slider.SliderThumbnail
}

func (self *Service) Init() {
	self.fileInPdf = fileutil.FileInPdf{}
	self.fileInPdf.Init()

	self.fileInWater = fileutil.FileInWater{}
	self.fileInWater.Init()

	self.fileOut = fileutil.FileOut{}
	self.fileOut.Init()

	self.selectRand = selectutil.SelectRandom{}
	self.selectCount = selectutil.SelectSingleton{}

	self.sliderClarity = slider.SliderClarity{}
	self.sliderClarity.Init()

	self.sliderThumbnail = slider.SliderThumbnail{}
	self.sliderThumbnail.Init()
}

func (self *Service) Run(myApp fyne.App, w fyne.Window) fyne.CanvasObject {
	fInPdf := self.fileInPdf.File(w)
	fInWater := self.fileInWater.File(w)
	fOut := self.fileOut.File(w)

	sRand := self.selectRand.SelectImageScreen(w)
	sCount := self.selectCount.SelectImageScreen(w)
	vc := container.New(layout.NewGridLayoutWithColumns(5), sRand, sCount, &widget.Label{Text: "水印图像缩略度"}, self.sliderThumbnail.Screen(w), &widget.Label{Text: "水印图像透明度"})
	bar := container.NewBorder(nil, nil, vc, nil, self.sliderClarity.Screen(w))
	pOut := progress.Screen(w)
	return container.New(layout.NewVBoxLayout(), fInPdf, fInWater, fOut, bar, pOut)
}
