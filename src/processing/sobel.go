package processing

import (
	"image"
	"image/color"
	"math"
)

var gx = [3][3]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}

var gy = [3][3]int{
	{-1, -2, -1},
	{0, 0, 0},
	{1, 2, 1},
}

func ApplySobel(img *image.Gray) *image.Gray {
	bounds := img.Bounds()
	result := image.NewGray(bounds)

	for y := 1; y < bounds.Max.Y-1; y++ {
		for x := 1; x < bounds.Max.X-1; x++ {
			var sumX, sumY int
			for j := -1; j <= 1; j++ {
				for i := -1; i <= 1; i++ {
					pixel := int(img.GrayAt(x+i, y+j).Y)
					sumX += pixel * gx[j+1][i+1]
					sumY += pixel * gy[j+1][i+1]
				}
			}
			val := uint8(math.Min(math.Sqrt(float64(sumX*sumX+sumY*sumY)), 255))
			result.SetGray(x, y, color.Gray{Y: val})
		}
	}
	return result
}
