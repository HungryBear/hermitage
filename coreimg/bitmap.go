package coreimg

import (
	"image"
	"image/png"
	"log"
	"os"
)

type Bitmap struct {
	Width, Height int
	Data          []RgbSpectrum
	Format        int
}

func NewBitmap(width, height int, format int) *Bitmap {
	return &Bitmap{Width: width, Height: height, Data: make([]RgbSpectrum, width*height), Format: format}
}

func (b *Bitmap) GetPixel(x, y int) *RgbSpectrum {
	return &b.Data[b.Width*y+x]
}

func (b *Bitmap) SetPixel(x, y int, spectrum *RgbSpectrum) {
	b.Data[b.Width*y+x] = *spectrum
}

func OpenImage(img image.Image) *Bitmap {
	var r, g, b, a uint32
	res := NewBitmap(img.Bounds().Max.X, img.Bounds().Max.Y, RgbFlagLinearRgb)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			color := img.At(x, y)
			r, g, b, a = color.RGBA()
			c := FromRgbInt(r, g, b, a)
			res.SetPixel(x, y, c)
		}
	}
	return res
}

func (b *Bitmap) SaveToImage() image.Image {
	rect := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: b.Width, Y: b.Height}}
	img := image.NewNRGBA(rect)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := b.GetPixel(x, y).ToRGB()
			img.Set(x, y, c)
			//color.NRGBA{R: uint8(255.0 * c.r), G: uint8(255.0 * c.g), B: uint8(255.0 * c.b), A: 255})
		}
	}
	return img
}

func (b *Bitmap) SaveToPng(filePath string) {
	img := b.SaveToImage()
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
