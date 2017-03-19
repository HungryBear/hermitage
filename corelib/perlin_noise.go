package corelib

import (
	"math"
	"math/rand"
)

const NoiseDim int = 15

var noise [NoiseDim + 1][NoiseDim + 1][NoiseDim + 1]float32

func init() {

	for i := 0; i < NoiseDim; i++ {
		for j := 0; j < NoiseDim; j++ {
			for k := 0; k < NoiseDim; k++ {
				noise[i][j][k] = rand.Float32()
			}
		}
	}
}

func mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}

func Spline(t float32) float32 {
	return t * t * (3.0 - 2.0*t)
}

func SineWave(t float32) float32 {
	return 0.5 * (1.0 + Sinf(t))
}

func PerlinNoise(x, y, z float32) float32 {
	sx := mod(x, float32(NoiseDim))
	sy := mod(y, float32(NoiseDim))
	sz := mod(z, float32(NoiseDim))

	ix := int(math.Trunc(float64(sx)))
	iy := int(math.Trunc(float64(sy)))
	iz := int(math.Trunc(float64(sz)))

	jx := ix + 1
	jy := iy + 1
	jz := iz + 1

	if jx >= NoiseDim {
		jx = 0
	}
	if jy >= NoiseDim {
		jy = 0
	}
	if jz >= NoiseDim {
		jz = 0
	}

	sx = Spline(sx - float32(ix))
	sy = Spline(sy - float32(iy))
	sz = Spline(sz - float32(iz))

	return (1-sx)*(1-sy)*(1-sz)*noise[ix][iy][iz] +
		(1-sx)*(1-sy)*sz*noise[ix][iy][jz] +
		(1-sx)*sy*(1-sz)*noise[ix][jy][iz] +
		(1-sx)*sy*sz*noise[ix][jy][jz] +
		sx*(1-sy)*(1-sz)*noise[jx][iy][iz] +
		sx*(1-sy)*sz*noise[jx][iy][jz] +
		sx*sy*(1-sz)*noise[jx][jy][iz] +
		sx*sy*sz*noise[jx][jy][jz]
}

func PerlinNoise3D(x, y, z float32) *Vector3 {
	sx := mod(x, float32(NoiseDim))
	sy := mod(y, float32(NoiseDim))
	sz := mod(z, float32(NoiseDim))

	ix := int(math.Trunc(float64(sx)))
	iy := int(math.Trunc(float64(sy)))
	iz := int(math.Trunc(float64(sz)))

	var jx, jy, jz int
	sx = Spline(sx - float32(ix))
	sy = Spline(sy - float32(iy))
	sz = Spline(sz - float32(iz))

	data := make([]float32, 3)

	for i := 0; i < 3; i++ {
		ix = (ix + 5) % NoiseDim
		iy = (iy + 5) % NoiseDim
		iz = (iz + 5) % NoiseDim

		jx = ix + 1
		jy = iy + 1
		jz = iz + 1

		if jx >= NoiseDim {
			jx = 0
		}
		if jy >= NoiseDim {
			jy = 0
		}
		if jz >= NoiseDim {
			jz = 0
		}
		data[i] = (1-sx)*(1-sy)*(1-sz)*noise[ix][iy][iz] +
			(1-sx)*(1-sy)*sz*noise[ix][iy][jz] +
			(1-sx)*sy*(1-sz)*noise[ix][jy][iz] +
			(1-sx)*sy*sz*noise[ix][jy][jz] +
			sx*(1-sy)*(1-sz)*noise[jx][iy][iz] +
			sx*(1-sy)*sz*noise[jx][iy][jz] +
			sx*sy*(1-sz)*noise[jx][jy][iz] +
			sx*sy*sz*noise[jx][jy][jz]
	}
	return NewVector3(data[0], data[1], data[2])
}

func Turbulence(p *Vector3, octaves int) float32 {
	var k float32 = 1.0
	var res float32 = 0.0
	var r Vector3 = *p.Copy()
	for i := 0; i < octaves; i++ {
		res += PerlinNoise(r.X, r.Y, r.Z) * k
		r = *r.Mulf(2.0)
		k *= 0.5
	}
	return res
}

func Turbulence3D(p *Vector3, octaves int) *Vector3 {
	var k float32 = 1.0
	var res Vector3 = *Unit()
	var r Vector3 = *p.Copy()
	for i := 0; i < octaves; i++ {
		res = *res.Add(PerlinNoise3D(r.X, r.Y, r.Z).Mulf(k))
		r = *r.Mulf(2.0)
		k *= 0.5
	}
	return &res
}
