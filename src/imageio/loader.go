package imageio

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func LoadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return image, nil
}
