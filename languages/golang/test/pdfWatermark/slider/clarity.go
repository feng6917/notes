package slider

import (
	"lgo/test/pdfWatermark/conf"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// 清晰度
type SliderClarity struct {
	Slider *widget.Slider
}

func (c *SliderClarity) Init() {
	c.Slider = widget.NewSlider(0, 100)
	c.Slider.Value = 100
	_ = conf.CF.SetSliderClarity(100)
}

func (c *SliderClarity) Screen(w fyne.Window) fyne.CanvasObject {
	c.Slider.OnChanged = func(val float64) {
		err := conf.CF.SetSliderClarity(val)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
	}
	return container.New(layout.NewGridLayout(1), c.Slider)
}
