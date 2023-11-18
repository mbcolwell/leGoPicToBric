package main

import (
	"os"

	picToBric "github.com/mbcolwell/leGoPicToBric/internal"
)

func main() {
	picToBric.ReadImage(os.Args[1])
}
