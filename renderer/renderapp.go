package main

//import "../corelib"
import (
	"../coreimg"
	"github.com/hungrybear/hermitage/corelib"
	"math/rand"
)

func main(){
	bmp:=coreimg.NewBitmap(320, 200, 1)

	for y := 0; y < 200; y++ {
		for x := 0; x < 320; x++ {
			c := corelib.PerlinNoise3D(rand.Float32(), rand.Float32(), rand.Float32())
			bmp.SetPixel(x, y, coreimg.NewRgbSpectrum(c.x, c.y, c.z, coreimg.RgbFlagLinearRgb))
		}
	}

	bmp.SaveToPng("D:\\generated.png")
}
