package imageutil

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"

	"github.com/nfnt/resize"
)

/*
* Scale 缩略图生成
* 入参:图片输入、输出，缩略图宽、高，精度
* 规则: 如果width 或 hight其中有一个为0，则大小不变 如果精度为0则精度保持不变
* 返回:缩略图真实宽、高、error
 */
func Scale(in io.Reader, out io.Writer, sliderThumbnail float64, quality int) (int, int, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	var (
		w, h int
	)
	origin, fm, err := image.Decode(in)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	width := int(float64(origin.Bounds().Max.X) * sliderThumbnail)
	height := int(float64(origin.Bounds().Max.Y) * sliderThumbnail)
	if quality == 0 {
		quality = 100
	}
	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)

	//return jpeg.Encode(out, canvas, &jpeg.Options{quality})
	w = canvas.Bounds().Dx()
	h = canvas.Bounds().Dy()
	switch fm {
	case "jpg", "JPG", "jpeg", "JPEG":
		return w, h, jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "png", "PNG":
		return w, h, png.Encode(out, canvas)
	case "gif":
		return w, h, gif.Encode(out, canvas, &gif.Options{})
	default:
		return w, h, errors.New("ERROR FORMAT")
	}
}
