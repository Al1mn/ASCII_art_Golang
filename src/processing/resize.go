package processing

import (
	"image"
	"math"
)

func Resize(img image.Image, newWidth int) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	ratio := float64(height) / float64(width)
	newHeight := int(math.Round(float64(newWidth) * ratio * 0.55))

	buff := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := int(float64(x) * float64(width) / float64(newWidth))
			srcY := int(float64(y) * float64(height) / float64(newHeight))
			buff.Set(x, y, img.At(srcX, srcY))
		}
	}
	return buff
}
