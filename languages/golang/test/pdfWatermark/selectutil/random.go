package selectutil

import (
	"lgo/test/pdfWatermark/conf"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
)

type SelectRandom struct {
	name    *widget.Select
	Current int
}

// selectImageScreen
func (c *SelectRandom) SelectImageScreen(_ fyne.Window) *fyne.Container {
	b := &SelectRandom{}
	nameList := []string{"固定水印位置", "随机水印位置"}
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
						conf.CF.SetRand(false)
						log.Infof("水印位置: %v", nameList[0])
					} else {
						conf.CF.SetRand(true)
						log.Infof("水印位置: %v", nameList[1])
					}
				}
				break
			}
		}
	})
	b.name.SetSelected(nameList[c.Current])

	return container.New(layout.NewVBoxLayout(), b.name)
}
