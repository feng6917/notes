package progress

import (
	"lgo/test/pdfWatermark/conf"
	"lgo/test/pdfWatermark/pdf"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	ProgressOut    *widget.ProgressBar
	infProgressOut *widget.ProgressBarInfinite
	endProgressOut chan interface{}
)

func startProgressObject() {
	ProgressOut.SetValue(0)
	select { // ignore stale end message
	case <-endProgressOut:
	default:
	}

	go func() {
		stopProgressObject()
	}()
	infProgressOut.Start()
}

func stopProgressObject() {
	if !infProgressOut.Running() {
		return
	}

	infProgressOut.Stop()
	endProgressOut <- struct{}{}
}

func Screen(w fyne.Window) fyne.CanvasObject {
	stopProgressObject()

	ProgressOut = widget.NewProgressBar()

	infProgressOut = widget.NewProgressBarInfinite()
	endProgressOut = make(chan interface{}, 1)
	startProgressObject()

	// 模拟一些操作后，切换到第二个对象
	switchToOutObject := widget.NewButton("输出文件", func() {
		log.Println("输出文件: ")
		// 1. 参数校验 必须指定类型 指定时间
		// if !CF.Download.IsDownloadOnline && !CF.Download.IsDownloadOffline && !CF.Download.IsDownloadHistory {
		// 	dialog.ShowError(errors.New("实时/离线/回溯 至少选择一个"), w)
		// 	return
		// }
		err := conf.CF.Check()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		// 参数校验完成设置进度 0.1
		ProgressOut.SetValue(0.1)
		// pdf.AddImageToPdf(conf.CF.InPdfPath, conf.CF.OutPath, conf.CF.InWaterPath, 0, 100, conf.CF.Scale, conf.CF.Random)
		// 初始化
		pdf := pdf.Pdf{}
		pdf.Init()
		ProgressOut.SetValue(0.13)

		err = pdf.InitImage(conf.CF.InWaterPath, conf.CF.SliderThumbnail)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		ProgressOut.SetValue(0.16)
		err = pdf.InitPdfReader(conf.CF.InPdfPath)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		ProgressOut.SetValue(0.20)
		ch := make(chan float64, 1)
		log.Info(conf.CF.Singleton, conf.CF.Random)
		pdf.Do(conf.CF.Singleton, conf.CF.Random, conf.CF.SliderClarity, ch)
		for {
			val, ok := <-ch
			if ok && val > 0 {
				ProgressOut.SetValue(0.20 + 0.6*val)
				if val == 1 {
					time.Sleep(time.Second * 2)
					break
				}
			}
		}
		err = pdf.Save(filepath.Join(conf.CF.OutPath, filepath.Base(conf.CF.InPdfPath)))
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		ProgressOut.SetValue(1)
		time.Sleep(1 * time.Second)
		w.Close()
	})

	return container.New(
		layout.NewFormLayout(),
		switchToOutObject,
		ProgressOut,
	)
}
