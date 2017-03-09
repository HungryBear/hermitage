package corelib

import (
	"testing"
	"math/rand"
	"github.com/hungrybear/hermitage/corelib"
)

func TestMersenneTwister_NextFloat(t *testing.T) {
	rnd:=NewMersenneTwister(rand.Int())
	for i:=0;i<1000;i++{
		c:=rnd.NextFloat()
		if c < 0 || c >= 1{
			t.Error("c is out of bound %f", c)
		}
	}
}

func TestMersenneTwister_Repeatable(t *testing.T) {
	seed:=rand.Int()
	rnd:=NewMersenneTwister(seed)
	elems:=make([]float32, 10)
	for i:=0;i<10;i++{
		elems[i]=rnd.NextFloat()
	}
	rnd = NewMersenneTwister(seed)
	for i:=0;i<10;i++ {
		c:=rnd.NextFloat()
		if !corelib.NearEqual(elems[i], c){
			t.Error("Elements from the seed %d are not equal enough %f %f", seed, elems[i],c)
		}
	}

}

// Benchmarking
func BenchmarkMersenneTwister_NextFloat(b *testing.B) {
	rnd:=NewMersenneTwister(rand.Int())
	var r float32
	for i:=0;i<b.N;i++{
		r+= rnd.NextFloat(i)
	}
}

func BenchmarkRandom(b*testing.B){
	var r float32
	for i:=0;i<b.N;i++{
		r+= rand.Float32()
	}
}

