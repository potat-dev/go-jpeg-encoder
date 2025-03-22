package main

import (
	"fmt"
	"os"

	"github.com/potat-dev/go-jpeg-encoder/internal/bmp"
)

func main() {
	file, err := os.Open("demo_image.bmp")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileHeader, err := bmp.ReadFileHeader(file)
	if err != nil {
		fmt.Println("Error reading file header:", err)
		return
	}

	infoHeader, err := bmp.ReadInfoHeader(file)
	if err != nil {
		fmt.Println("Error reading info header:", err)
		return
	}

	bmp.PrintFileHeader(fileHeader)
	bmp.PrintInfoHeader(infoHeader)
}
