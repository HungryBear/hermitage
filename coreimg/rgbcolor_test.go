package coreimg

import (
	"testing"
	"math/rand"
	"hermitage/corelib"
)

func TestRgbColorConversion(t*testing.T) {
	col := NewRgbSpectrum(rand.Float32(), rand.Float32(), rand.Float32(), RgbFlagLinearRgb)
	xyz := SRGB.ConvertToXyz(col.r, col.g, col.b)

	if !corelib.HasFlag(xyz.flags, RgbFlagsXyz){
		t.Error("Invalid color flag - should contain XYZ")
	}

	rev_col := SRGB.ConvertToRgb(xyz.r, xyz.g, xyz.b);

	if !corelib.NearEqualEps(col.r, rev_col.r, 0.001){
		t.Error("Invalid r component after conversion " , col.r, rev_col.r)
	}
	if !corelib.NearEqualEps(col.g, rev_col.g, 0.001){
		t.Error("Invalid g component after conversion " , col.g, rev_col.g)
	}
	if !corelib.NearEqualEps(col.b, rev_col.b, 0.001){
		t.Error("Invalid b component after conversion " , col.b, rev_col.b)
	}
}

func TestRgbColorArithmetic(t*testing.T){
	x := NewRgbSpectrum(rand.Float32(), rand.Float32(), rand.Float32(), RgbFlagLinearRgb)
	y := NewRgbSpectrum(rand.Float32(), rand.Float32(), rand.Float32(), RgbFlagLinearRgb)

	x_s := x

	x.Mul(y)
	x.Mul(NewRgbSpectrum(1.0/y.r, 1.0/y.g, 1.0/y.b, RgbFlagLinearRgb))

	x.Mulf(2.0)
	x.Mulf(0.5)

	if !corelib.NearEqual(x.r, x_s.r){
		t.Error("Arithmetic error ", x.r, x_s.r)
	}

	if !corelib.NearEqual(x.g, x_s.g){
		t.Error("Arithmetic error ", x.g, x_s.g)
	}

	if !corelib.NearEqual(x.b, x_s.b){
		t.Error("Arithmetic error ", x.b, x_s.b)
	}
}
