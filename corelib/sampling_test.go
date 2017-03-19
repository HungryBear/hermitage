package corelib

import (
	"math/rand"
	"testing"
)

func TestMersenneTwister_NextFloat(t *testing.T) {
	rnd := NewMersenneTwister(rand.Int())
	for i := 0; i < 1000; i++ {
		c := rnd.NextFloat()
		if c < 0 || c >= 1 {
			t.Error("c is out of bound %f", c)
		}
	}
}

func TestMersenneTwister_NextFloatUniformity(t *testing.T) {
	rnd := NewMersenneTwister(rand.Int())
	rnd.Initialize(rand.Uint32())
	counter := 0
	for i := 0; i < 1000; i++ {
		c := rnd.NextFloat()
		if c < 0.5 {
			counter++
		}
	}
	if counter > 550 {
		t.Error("uniformity is broken - from 1000 iterations %d < 0.5", counter)
	}
}

func TestLcgRandom_NextFloat(t *testing.T) {
	rnd := NewLcgRandom(rand.Int())
	rnd.Initialize(rand.Uint32())
	for i := 0; i < 1000; i++ {
		c := rnd.NextFloat()
		if c < 0 || c >= 1 {
			t.Error("c is out of bound %f", c)
		}
	}
}

func TestMersenneTwister_Repeatable(t *testing.T) {
	seed := rand.Int()
	rnd := NewMersenneTwister(seed)
	elems := make([]float32, 10)
	for i := 0; i < 10; i++ {
		elems[i] = rnd.NextFloat()
	}
	rnd = NewMersenneTwister(seed)
	for i := 0; i < 10; i++ {
		c := rnd.NextFloat()
		if !NearEqual(elems[i], c) {
			t.Error("Elements from the seed %d are not equal enough %f %f", seed, elems[i], c)
		}
	}
}

func TestNewLcgRandom_NextFloatUniformity(t *testing.T) {
	rnd := NewLcgRandom(rand.Int())
	counter := 0
	for i := 0; i < 1000; i++ {
		c := rnd.NextFloat()
		if c < 0.5 {
			counter++
		}
	}
	if counter > 575 { //TODO fix it
		t.Error("uniformity is broken - from 1000 iterations %d < 0.5", counter)
	}
}

func TestNewSamplesCollection(t *testing.T) {
	sampler := NewUniformSampler(120, 100, rand.Int(), NewMersenneTwister(rand.Int()))
	sc := sampler.RequestSamples(10, 10)
	if sc.Samples1DCount != 10 && sc.Samples2DCount != 10 {
		t.Error("Invalid samples count")
	}
	for i := 0; i < 10; i++ {
		if sc.Samples1D[i] < Epsilon {
			t.Error("Invalid 1D sample value %f", sc.Samples1D[i])
		}
		if sc.Samples2D[i].X < Epsilon {
			t.Error("Invalid 2D sample x value %f", sc.Samples2D[i].X)
		}
		if sc.Samples2D[i].Y < Epsilon {
			t.Error("Invalid 2D sample y value %f", sc.Samples2D[i].Y)
		}
	}

}

// Benchmarking
func BenchmarkMersenneTwister_NextFloat(b *testing.B) {
	rnd := NewMersenneTwister(rand.Int())
	var r float32
	for i := 0; i < b.N; i++ {
		r += rnd.NextFloat(i)
	}
}

func BenchmarkRandom(b *testing.B) {
	var r float32
	for i := 0; i < b.N; i++ {
		r += rand.Float32()
	}
}
