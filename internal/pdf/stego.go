package pdf

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
)

func EmbedPayloadInPNG(imgPath string, payload []byte) ([]byte, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	bits := bytesToBits(payload)
	bitIndex := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8, g8, b8, a8 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)

			if bitIndex < len(bits) {
				r8 = (r8 & 0xFE) | bits[bitIndex]
				bitIndex++
			}
			if bitIndex < len(bits) {
				g8 = (g8 & 0xFE) | bits[bitIndex]
				bitIndex++
			}
			if bitIndex < len(bits) {
				b8 = (b8 & 0xFE) | bits[bitIndex]
				bitIndex++
			}
			newImg.Set(x, y, color.RGBA{r8, g8, b8, a8})
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, newImg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ExtractPayloadFromPNG(imgPath string) ([]byte, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	var bits []byte

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
			bits = append(bits, r8&1, g8&1, b8&1)
		}
	}
	return bitsToBytes(bits), nil
}

func bytesToBits(data []byte) []byte {
	bits := make([]byte, len(data)*8)
	for i, b := range data {
		for j := 0; j < 8; j++ {
			bits[i*8+j] = (b >> (7 - j)) & 1
		}
	}
	return bits
}

func bitsToBytes(bits []byte) []byte {
	data := make([]byte, len(bits)/8)
	for i := 0; i < len(data); i++ {
		var b byte
		for j := 0; j < 8; j++ {
			b |= bits[i*8+j] << (7 - j)
		}
		data[i] = b
	}
	return data
}