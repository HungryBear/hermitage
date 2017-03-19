package corelib

import (
	"testing"
	"math/rand"
)

func randMatrix()*Matrix4x4{
	return NewMatrix4x4(rand.Float32(), rand.Float32(),rand.Float32(),rand.Float32(),
		rand.Float32(),rand.Float32(),rand.Float32(),rand.Float32(),
		rand.Float32(),rand.Float32(),rand.Float32(),rand.Float32(),
		rand.Float32(),rand.Float32(),rand.Float32(),rand.Float32())
}

func Test_Matrix4x4Arithmetic(t*testing.T){
	m1:=randMatrix()
	m2:=m1.Mul(Identity())

	if !NearEqual(m1[0][0], m2[0][0]){
		t.Errorf("invalid identity multiplication %f %f",m1[0][0], m2[0][0])
	}

}
