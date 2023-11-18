package picToBric

import (
	"fmt"
	"image"
	"log"
	"os"
	_ "image/png"
	_ "image/jpeg"
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

	fmt.Println(imageData)
	fmt.Println(imageType)

}
