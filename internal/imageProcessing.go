package picToBric

import (
	"fmt"
	"image"

	// "image/color"
	"image/jpeg"
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

	pixels := imgToPixels(imageData)
	clusters := kmeans(10, pixels)

	f, err := os.Create("pixel-test.jpg")
	if err != nil {
		log.Fatalf("error creating file: %s", err)
	}
	defer f.Close()

	err = jpeg.Encode(f, clusteredPixelsImg(pixels, clusters), nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
