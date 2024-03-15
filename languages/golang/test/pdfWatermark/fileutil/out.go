package fileutil

import (
	"lgo/test/pdfWatermark/conf"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type FileOut struct {
	Entry *widget.Entry
}

func (c *FileOut) Init() {
	entry := widget.NewEntry()
	str, _ := os.Getwd()
	entry.SetPlaceHolder(str)
	c.Entry = entry
}

func (c *FileOut) File(w fyne.Window) fyne.CanvasObject {
	openFolder := widget.NewButton("选择PDF文件输出地址", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if list == nil {
				log.Println("Cancelled")
				return
			}

			fileDir := strings.Replace(list.String(), "file://", "", 1)
			c.Entry.SetText(fileDir)
			err = conf.CF.SetOut(fileDir)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			// c.dirName = fileDir
			log.Println("当前PDF文件路径: ", fileDir)
		}, w)
	})

	return container.New(layout.NewFormLayout(), openFolder, c.Entry)
}
