package coreimg

import (
	"testing"
)

func TestRgbColorConversion(t*testing.T) {

	col := NewRgbSpectrum(1.0, 0.4, 0.6, RgbFlagLinearRgb)

	converted := col.Convert(SRGB)

	if (converted.r == col.r){
		t.Error("Impossibro")
	}

}
