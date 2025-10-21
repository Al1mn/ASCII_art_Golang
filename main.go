package main

import (
	"console_project/src/ascii"
	"console_project/src/imageio"
	"console_project/src/processing"
	"fmt"
	"log"
)

func main() {
	filePath := "Pictures/Tima.jpeg"

	img, err := imageio.LoadImage(filePath)
	if err != nil {
		log.Fatal(err)
	}

	resized := processing.Resize(img, 500)
	grayScale := processing.ToLuma(resized)
	sobel := processing.ApplySobel(grayScale)

	fmt.Println(ascii.ConvertToASCII(sobel, resized, true))
}
