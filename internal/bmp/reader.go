package bmp

import (
	"encoding/binary"
	"fmt"
	"os"
)

// ReadFileHeader reads the BMP file header from a file
func ReadFileHeader(file *os.File) (*FileHeader, error) {
	var header FileHeader
	err := binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}
	return &header, nil
}

// ReadInfoHeader reads the BMP info header from a file
func ReadInfoHeader(file *os.File) (*InfoHeader, error) {
	var infoHeader InfoHeader
	err := binary.Read(file, binary.LittleEndian, &infoHeader)
	if err != nil {
		return nil, err
	}
	return &infoHeader, nil
}

// PrintFileHeader prints the BMP file header in a readable format
func PrintFileHeader(header *FileHeader) {
	fmt.Println("======= BMP File Header =======")
	fmt.Printf("File Type:          %s\n", string(header.FileType[:]))
	fmt.Printf("File Size:          %d bytes\n", header.FileSize)
	fmt.Printf("Reserved:           0x%X\n", header.Reserved)
	fmt.Printf("Data Offset:        %d bytes\n", header.DataOffset)
	fmt.Println("===============================")
}

// PrintInfoHeader prints the BMP info header in a readable format
func PrintInfoHeader(infoHeader *InfoHeader) {
	fmt.Println("======= BMP Info Header =======")
	fmt.Printf("Header Size:        %d bytes\n", infoHeader.HeaderSize)
	fmt.Printf("Image Width:        %d px\n", infoHeader.Width)
	fmt.Printf("Image Height:       %d px\n", infoHeader.Height)
	fmt.Printf("Color Planes:       %d\n", infoHeader.Planes)
	fmt.Printf("Bits Per Pixel:     %d\n", infoHeader.BitsPerPixel)
	fmt.Printf("Compression Method: %d\n", infoHeader.Compression)
	fmt.Printf("Image Size:         %d bytes\n", infoHeader.ImageSize)
	fmt.Printf("X Resolution:       %d px/m\n", infoHeader.XPixelsPerMeter)
	fmt.Printf("Y Resolution:       %d px/m\n", infoHeader.YPixelsPerMeter)
	fmt.Printf("Colors Used:        %d\n", infoHeader.ColorsUsed)
	fmt.Printf("Important Colors:   %d\n", infoHeader.ImportantColors)
	fmt.Println("===============================")
}
