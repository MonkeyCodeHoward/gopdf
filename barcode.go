package gopdf

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"image/png"
	"os"
)

func Encode(content, savePath string, width, height int) error {
	enc := oned.NewCode128Writer()
	img, err := enc.Encode(content, gozxing.BarcodeFormat_CODE_128, width, height, nil)
	if err != nil {
		return err
	}
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}
