package corelib

import (
	"math"
	"math/rand"
	"testing"
)

func TestClamp(t *testing.T) {
	iterations := rand.Intn(1000)
	for i := 0; i < iterations; i++ {
		c := Clamp(2.0*rand.Float32()-1, 0, 1.0)
		if c > 1 || c < 0 {
			t.Error("Invalid value after clamp ", c)
		}
	}
}

func TestSqrtf(t *testing.T) {
	iterations := rand.Intn(100)
	for i := 0; i < iterations; i++ {
		angle := float32(rand.Intn(100)) * rand.Float32()
		goCos := float32(math.Sqrt(float64(angle)))
		libCos := Sqrtf(angle)

		if !NearEqualEps(goCos, libCos, 0.9) {
			t.Error("Sqrt error [go / lib]  - value ", goCos, libCos, angle)
		}
	}
}

func TestLerp(t *testing.T) {
	iterations := rand.Intn(100)
	for i := 0; i < iterations; i++ {
		scale := rand.Float32()
		t0 := float32(rand.Intn(10)) * scale
		t1 := float32(10.0+rand.Intn(10)) * scale

		if c := Lerp(0.5, t0, t1); c < t0 || c > t1 {
			t.Error("Linear interpolation error [t0..t1] t", t0, t1, c)
		}
	}
}
