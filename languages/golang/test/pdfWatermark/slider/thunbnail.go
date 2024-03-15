package slider

import (
	"lgo/test/pdfWatermark/conf"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// 缩略度
type SliderThumbnail struct {
	Slider *widget.Slider
}

func (c *SliderThumbnail) Init() {
	c.Slider = widget.NewSlider(0, 100)
	c.Slider.Value = 100
	_ = conf.CF.SetSliderThumbnail(100)
}

func (c *SliderThumbnail) Screen(w fyne.Window) fyne.CanvasObject {
	c.Slider.OnChanged = func(val float64) {
		err := conf.CF.SetSliderThumbnail(val)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
	}
	return container.New(layout.NewGridLayout(1), c.Slider)
}
