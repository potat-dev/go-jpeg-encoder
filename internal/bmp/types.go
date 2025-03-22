package bmp

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
