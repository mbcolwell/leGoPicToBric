package picToBric

import (
	"fmt"
	"image"

	// "image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func ReadImage(path string) {
	imageFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file")
	}
	defer imageFile.Close()

	imageData, imageType, err := image.Decode(imageFile)
	if err != nil {
		log.Fatalf("Unable to decode file contents into image data (%v)", err)
	}

	fmt.Println(imageData.Bounds().Size())
	fmt.Println(imageType)
	r, g, b, a := imageData.At(0, 1).RGBA()
	fmt.Println(r >> 8)
	fmt.Println(g >> 8)
	fmt.Println(b >> 8)
	fmt.Println(a >> 8)

	fmt.Println(imgToPixels(imageData)[0][1])

}
