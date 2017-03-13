package coreimg

import (
	"github.com/hungrybear/hermitage/corelib"
	"testing"
)

func TestBitmap(t *testing.T) {
	bmp := NewBitmap(120, 100, 1)

	for y := 0; y < 100; y++ {
		for x := 0; x < 120; x++ {
			bmp.SetPixel(x, y, FromRgbInt(127, 32, 74, 255))
		}
	}

	pix := bmp.GetPixel(60, 50)

	if pix.IsBlack() {
		t.Error("Invalid pixel value in the center")
	}

	img := bmp.SaveToImage()

	newBmp := OpenImage(img)

	newPix := newBmp.GetPixel(60, 50)

	if !corelib.NearEqual(pix.r, newPix.r) {
		t.Error("invalid value after conversion", pix.r, newPix.r)
	}

}
