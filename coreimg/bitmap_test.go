package coreimg

import (
	"github.com/hungrybear/hermitage/corelib"
	"testing"
)

func TestBitmap(t *testing.T) {
	bmp := NewBitmap(120, 100, 1)

	for y := 0; y < 100; y++ {
		for x := 0; x < 120; x++ {
			bmp.SetPixel(x, y, FromRgbInt(127+uint32(y), 32, uint32(x+20), 255))
		}
	}

	pix := bmp.GetPixel(60, 50)

	t.Logf("%f %f %f ", pix.r, pix.g, pix.b)



	if pix.IsBlack() {
		t.Error("Invalid pixel value in the center")
	}

	img := bmp.SaveToImage()

	bmp.SaveToPng("D:\\test.png")

	var r,g, b, a uint32
	r,g,b,a = img.At(60, 50).RGBA()

	t.Logf("Int pixel %d %d %d %d", r,g,b, a)
	newBmp := OpenImage(img)

	newPix := newBmp.GetPixel(60, 50)

	if !corelib.NearEqualEps(pix.r, newPix.r/255.0, 0.01) {
		t.Error("invalid value after conversion", pix.r, newPix.r/255.0)
	}

}
