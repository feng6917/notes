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
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type FileInPdf struct {
	Entry *widget.Entry
}

func (c *FileInPdf) Init() {
	entry := widget.NewEntry()
	str, _ := os.Getwd()
	entry.SetPlaceHolder(str)
	c.Entry = entry
}

func (c *FileInPdf) File(w fyne.Window) fyne.CanvasObject {
	openFolder := widget.NewButton("选择PDF文件地址", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader == nil {
				return
			}
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			// fileOpened(reader)
			fileDir := strings.Replace(reader.URI().String(), "file://", "", 1)
			c.Entry.SetText(fileDir)
			// c.fileName = fileDir
			err = conf.CF.SetInPdf(fileDir)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			log.Println("当前PDF文件: ", fileDir)
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf", ".PDF"}))
		fd.Show()
	})

	return container.New(layout.NewFormLayout(), openFolder, c.Entry)
}
