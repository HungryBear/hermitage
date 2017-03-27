package main

import (
	"../coreimg"
	"../corelib"
	"../scenelib"
)

var rough float32 = 0.92
var bitmap coreimg.Bitmap = *coreimg.NewBitmap(512, 512, 1)

func main(){

	wc := scenelib.Scene{}
	wc.Name = "New Scene"
	hm:=corelib.GenerateDSQ(512, 512, rough)

	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			bitmap.SetPixel(x, y, coreimg.CreateRgbSpectrum(hm[x + y*512], 1))
		}
	}


	bitmap.SaveToPng("D:\\dsq_2.png")
}
