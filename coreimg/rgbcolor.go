package coreimg

import (
	"github.com/hungrybear/hermitage/corelib"
	"image/color"
	"math"
)

const (
	RgbFlagUnknown = 1 << iota
	RgbFlagLinearRgb
	RgbFlagsRgb
	RgbFlagsXyz
	RgbFlagCustom
	RgbFlagGammaCorrected
)

type RgbConvertible interface {
	Convert(space ColorSpace) *RgbSpectrum
}

type RgbConvertor interface {
	ConvertToRgb(x, y, z float32) *RgbSpectrum
	ConvertToXyz(r, g, b float32) *RgbSpectrum
}

type RgbSpectrum struct {
	r, g, b float32
	flags   int
}

type ColorSpace struct {
	RWeight, GWeight, BWeight [3]float32
	XWeight, YWeight, ZWeight [3]float32
	Xw, Yw, Zw                float32
}

var SRGB ColorSpace = ColorSpace{
	XWeight: [3]float32{0.4124564, 0.3575761, 0.1804375},
	YWeight: [3]float32{0.2126729, 0.7151522, 0.0721750},
	ZWeight: [3]float32{0.0193339, 0.1191920, 0.9503041},
	RWeight: [3]float32{3.2404542, -1.5371385, -0.4985314},
	GWeight: [3]float32{-0.9692660, 1.8760108, 0.0415560},
	BWeight: [3]float32{0.0556434, -0.2040259, 1.0572252},
	Xw:      0.95047, Yw: 1.0, Zw: 1.08883} //D65 + sRGB

func NewRgbSpectrum(red, green, blue float32, flags int) *RgbSpectrum {
	return &RgbSpectrum{r: red, g: green, b: blue, flags: flags}
}

func CreateRgbSpectrum(a float32, flags int) *RgbSpectrum {
	return &RgbSpectrum{r: a, g: a, b: a, flags: flags}
}

func (r *RgbSpectrum) Add(s *RgbSpectrum) *RgbSpectrum {
	r.r += s.r
	r.g += s.g
	r.b += s.b
	return r
}

func (r *RgbSpectrum) Mul(s *RgbSpectrum) *RgbSpectrum {
	r.r *= s.r
	r.g *= s.g
	r.b *= s.b
	return r
}

func (r *RgbSpectrum) Mulf(s float32) *RgbSpectrum {
	r.r *= s
	r.g *= s
	r.b *= s
	return r
}

func (r *RgbSpectrum) IsBlack() bool {
	return corelib.NearEqual(r.r, 0) && corelib.NearEqual(r.g, 0) && corelib.NearEqual(r.b, 0)
}

func FromRgbInt(r, g, b, a uint32) *RgbSpectrum {
	return NewRgbSpectrum(float32(r)/255.0, float32(g)/255.0, float32(b)/255.0, RgbFlagLinearRgb)
}

func (r *RgbSpectrum) ToRGB() color.NRGBA {
	return color.NRGBA{
		R: uint8(math.Min(float64(r.r*255.0), 255.0)),
		G: uint8(math.Min(float64(r.g*255.0), 255.0)),
		B: uint8(math.Min(float64(r.b*255.0), 255.0)),
		A:255}
}

func (r *RgbSpectrum) Convert(space ColorSpace) *RgbSpectrum {
	return space.ConvertToRgb(r.r, r.g, r.b)
}

func (c *ColorSpace) ConvertToRgb(x, y, z float32) *RgbSpectrum {
	return &RgbSpectrum{
		flags: RgbFlagsRgb,
		r:     c.RWeight[0]*x + c.RWeight[1]*y + c.RWeight[2]*z,
		g:     c.GWeight[0]*x + c.GWeight[1]*y + c.GWeight[2]*z,
		b:     c.BWeight[0]*x + c.BWeight[1]*y + c.BWeight[2]*z}
}

func (c *ColorSpace) ConvertToXyz(x, y, z float32) *RgbSpectrum {
	return &RgbSpectrum{
		flags: RgbFlagsXyz,
		r:     c.XWeight[0]*x + c.XWeight[1]*y + c.XWeight[2]*z,
		g:     c.YWeight[0]*x + c.YWeight[1]*y + c.YWeight[2]*z,
		b:     c.ZWeight[0]*x + c.ZWeight[1]*y + c.ZWeight[2]*z}
}
