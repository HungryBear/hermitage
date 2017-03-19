package coreimg

import (
	"hermitage/corelib"
	"math/rand"
	"testing"
)

func TestRgbColorConversion(t *testing.T) {
	col := NewRgbSpectrum(rand.Float32(), rand.Float32(), rand.Float32(), RgbFlagLinearRgb)
	xyz := SRGB.ConvertToXyz(col.R, col.G, col.B)

	if !corelib.HasFlag(xyz.Flags, RgbFlagsXyz) {
		t.Error("Invalid color flag - should contain XYZ")
	}

	rev_col := SRGB.ConvertToRgb(xyz.R, xyz.G, xyz.B)

	if !corelib.NearEqualEps(col.R, rev_col.R, 0.001) {
		t.Error("Invalid r component after conversion ", col.R, rev_col.R)
	}
	if !corelib.NearEqualEps(col.G, rev_col.G, 0.001) {
		t.Error("Invalid g component after conversion ", col.G, rev_col.G)
	}
	if !corelib.NearEqualEps(col.B, rev_col.B, 0.001) {
		t.Error("Invalid b component after conversion ", col.B, rev_col.B)
	}
}

func TestRgbColorArithmetic(t *testing.T) {
	x := NewRgbSpectrum(rand.Float32(), rand.Float32(), rand.Float32(), RgbFlagLinearRgb)
	y := NewRgbSpectrum(rand.Float32(), rand.Float32(), rand.Float32(), RgbFlagLinearRgb)

	x_s := x

	x.Mul(y)
	x.Mul(NewRgbSpectrum(1.0/y.R, 1.0/y.G, 1.0/y.B, RgbFlagLinearRgb))

	x.Mulf(2.0)
	x.Mulf(0.5)

	if !corelib.NearEqual(x.R, x_s.R) {
		t.Error("Arithmetic error ", x.R, x_s.R)
	}

	if !corelib.NearEqual(x.G, x_s.G) {
		t.Error("Arithmetic error ", x.G, x_s.G)
	}

	if !corelib.NearEqual(x.B, x_s.B) {
		t.Error("Arithmetic error ", x.B, x_s.B)
	}
}
