package bmp

import (
	"encoding/binary"
	"fmt"
	"os"
)

// FileHeader represents the 14-byte BMP file header
type FileHeader struct {
	FileType   [2]byte // 2 bytes ("BM" for Bitmap)
	FileSize   uint32  // 4 bytes (Size of file)
	Reserved   uint32  // 4 bytes (Reserved, typically 0)
	DataOffset uint32  // 4 bytes (Offset where pixel array starts)
}

// InfoHeader represents the 40-byte DIB header (BITMAPINFOHEADER)
type InfoHeader struct {
	HeaderSize      uint32 // 4 bytes (Size of this header)
	Width           int32  // 4 bytes (Bitmap width in pixels)
	Height          int32  // 4 bytes (Bitmap height in pixels)
	Planes          uint16 // 2 bytes (Number of color planes, must be 1)
	BitsPerPixel    uint16 // 2 bytes (Bits per pixel, e.g., 24 for RGB)
	Compression     uint32 // 4 bytes (Compression method)
	ImageSize       uint32 // 4 bytes (Size of raw bitmap data)
	XPixelsPerMeter int32  // 4 bytes (Horizontal resolution, pixels/meter)
	YPixelsPerMeter int32  // 4 bytes (Vertical resolution, pixels/meter)
	ColorsUsed      uint32 // 4 bytes (Number of colors in palette)
	ImportantColors uint32 // 4 bytes (Important colors used)
}

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
