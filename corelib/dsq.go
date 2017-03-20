package corelib

import "math/rand"

func diamondStep(sx, sy, rx, ry, w int, bitmap []float32, rough float32){
	dx:=(sx + rx) / 2
	dy:=(sy + ry) / 2

	c0:=bitmap[dx+ sy*w]
	c1:=bitmap[dx+ ry*w]
	c2:=bitmap[rx+ dy*w]
	c3:=bitmap[rx+ dy*w]

	c:=(c0+c1+c2+c3)*0.25+(1.0 / (1.0+ rand.Float32()*rough));

	bitmap[dx+dy*w] = c
}

func squareStep(sx, sy, rx, ry, w, h int, bitmap []float32, rough float32){
	lx:=(sx+rx)/ 2
	ly:=(sy+ry)/ 2

	c0:=bitmap[sx+ sy*w]
	c1:=bitmap[sx+ ry*w]
	c2:=bitmap[rx+ sy*w]
	c3:=bitmap[rx+ ry*w]

	if lx >= w{
		lx = w - lx
	}
	if ly >= h{
		ly = h - ly
	}
	c:=(c0+c1+c2+c3)*0.25+(rand.Float32()*rough);

	bitmap[lx+ly*w] = c
}


func diamondSquare(rx, ry,lx, ly,w, h int, bitmap []float32, rough float32){
	xl:=lx-rx
	yl:=ly-ry
	if xl<=1 || yl <= 1 {
		return
	}
	diamondStep(rx, ry, lx, ly,w, bitmap,rough)
	squareStep(rx, ry, lx, ly, w, h, bitmap, rough)

	x0 :=(rx + lx )/ 2
	y0 :=(ry + ly )/ 2

	diamondSquare(rx, ry, x0, y0, w, h, bitmap, rough)
	diamondSquare(rx, y0, x0, ly, w, h, bitmap, rough)
	diamondSquare(x0, y0, lx, ly, w, h, bitmap, rough)
	diamondSquare(x0, ry, lx, y0, w, h, bitmap, rough)
}

func GenerateDSQ(width, height int, rough float32) []float32{
	bitmap:=make([]float32, width*height)
	bitmap[0] = rand.Float32()
	bitmap[(height-1)*width] = rand.Float32()
	bitmap[width-1] = rand.Float32()
	bitmap[width-1 + (height-1)*width] = rand.Float32()
	diamondSquare(0, 0, width-1, height-1, width, height, bitmap, rough)
	return bitmap
}
