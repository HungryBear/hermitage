package corelib

import "testing"
import "math/rand"

func TestHaltonSequence_NextFloat(t *testing.T) {
	rnd := NewHaltonSequence(1, rand.Int())
	for i := 0; i < 1000; i++ {
		c := rnd.NextFloat()
		if c < 0 || c >= 1 {
			t.Error("c is out of bound %f", c)
		}
	}
}

func TestHaltonSequence_NextFloatUniformity(t *testing.T) {
	rnd := NewHaltonSequence(1, rand.Int())
	rnd.Initialize(rand.Uint32())
	counter := 0
	for i := 0; i < 1000; i++ {
		c := rnd.NextFloat()
		if c < 0.5 {
			counter++
		}
	}
	if counter > 500 {
		t.Error("uniformity is really broken - from 1000 iterations %d < 0.5", counter)
	}
	t.Log("values < 0.5 count = ", counter)
}
