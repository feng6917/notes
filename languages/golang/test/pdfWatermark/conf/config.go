package conf

import (
	"errors"
	"lgo/util/file"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Conf struct {
	InPdfPath       string
	InWaterPath     string
	OutPath         string
	Random          bool
	Singleton       bool    // 单张
	SliderClarity   float64 // 清晰度
	SliderThumbnail float64 // 缩略尺度
}

var CF Conf

func (self *Conf) SetInPdf(val string) error {
	_, err := os.Stat(val)
	if err != nil {
		log.Errorf("filepath not found, err: %v", err.Error())
		return errors.New("file not found.")
	}
	vd := strings.ToLower(val)
	if !strings.Contains(vd, "pdf") {
		return errors.New("file not pdf style.")
	}
	self.InPdfPath = val
	return err
}

func (self *Conf) SetInWater(val string) error {
	_, err := os.Stat(val)
	if err != nil {
		log.Errorf("filepath not found, err: %v", err.Error())
		return errors.New("file not found.")
	}
	vd := strings.ToLower(val)
	if !strings.Contains(vd, "png") && !strings.Contains(vd, "jpg") && !strings.Contains(vd, "jpeg") {
		return errors.New("file not pdf style.")
	}
	self.InWaterPath = val
	return err
}

func (self *Conf) SetOut(val string) error {
	_, err := os.Stat(val)
	if err != nil {
		log.Errorf("filepath not found, err: %v", err.Error())
		return errors.New("file not found.")
	}
	self.OutPath = val
	return err
}

func (self *Conf) SetRand(val bool) {
	self.Random = val
}

func (self *Conf) SetSingleton(val bool) {
	self.Singleton = val
}

func (self *Conf) SetSliderClarity(val float64) error {
	log.Info("水印图像透明度: ", val/100)
	self.SliderClarity = val / 100
	return nil
}

func (self *Conf) SetSliderThumbnail(val float64) error {
	log.Info("水印图像缩略度: ", val/100)
	self.SliderThumbnail = val / 100
	return nil
}

func (self *Conf) Check() error {
	if time.Now().Year() > 2024 {
		return errors.New("软件异常,请联系开发人员!")
	}
	// check path exist
	if !file.PathExists(self.InPdfPath) {
		return errors.New("in pdf path not exist!")
	}

	if !file.PathExists(self.InWaterPath) {
		return errors.New("in water path not exist!")
	}

	if !file.PathExists(self.OutPath) {
		return errors.New("out path not exist!")
	}
	if self.SliderClarity <= 0.05 {
		return errors.New("透明度 过小!")
	}
	if self.SliderThumbnail <= 0.05 {
		return errors.New("缩略度 过小!")
	}
	return nil
}
