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

type FileInWater struct {
	Entry    *widget.Entry
}

func (c *FileInWater) Init() {
	entry := widget.NewEntry()
	str, _ := os.Getwd()
	entry.SetPlaceHolder(str)
	c.Entry = entry
}

func (c *FileInWater) File(w fyne.Window) fyne.CanvasObject {
	openFolder := widget.NewButton("选择水印文件地址", func() {
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
			err = conf.CF.SetInWater(fileDir)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			log.Println("当前水印文件路径: ", fileDir)
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".PNG", ".jpg", ".JPG", ".jpeg", ".JPEG"}))
		fd.Show()
	})

	return container.New(layout.NewFormLayout(), openFolder, c.Entry)
}

func fileOpened(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}

	ext := f.URI().Extension()

	if strings.ToLower(ext) == ".pdf" {
		log.Infof("pdf file path: %s", f.URI().String())
	}
	err := f.Close()
	if err != nil {
		fyne.LogError("Failed to close stream", err)
	}
}
