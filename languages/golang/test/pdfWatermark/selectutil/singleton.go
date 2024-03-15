package selectutil

import (
	"lgo/test/pdfWatermark/conf"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
)

type SelectSingleton struct {
	name    *widget.Select
	Current int
}

// selectImageScreen
func (c *SelectSingleton) SelectImageScreen(_ fyne.Window) *fyne.Container {
	b := &SelectRandom{}
	nameList := []string{"单张水印图像", "多张水印图像"}
	b.name = widget.NewSelect(nameList, func(name string) {
		for i, name := range nameList {
			if name == b.name.Selected {
				if c.Current != i {
					if i < 0 || i > len(nameList)-1 {
						return
					}
					c.Current = i
					b.name.SetSelected(nameList[i])
					if i == 0 {
						conf.CF.SetSingleton(true)
						log.Infof("水印图像数量: %v", nameList[0])
					} else {
						conf.CF.SetSingleton(false)
						log.Infof("水印图像数量: %v", nameList[1])
					}
				}
				break
			}
		}
	})
	b.name.SetSelected(nameList[c.Current])
	conf.CF.SetSingleton(true)

	return container.New(layout.NewVBoxLayout(), b.name)
}
