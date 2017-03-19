package corelib

import (
	"math/rand"
	"testing"
)

func randV3() *Vector3 {
	return NewVector3(rand.Float32(), rand.Float32(), rand.Float32())
}

func randV2() *Vector2 {
	return NewVector2(rand.Float32(), rand.Float32())
}

func randV4() *Vector4 {
	return NewVector4(rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32())
}

func TestOnb_Ops(t *testing.T) {
	onb := NewOnb(*randV3())
	local := randV3()
	world := onb.ToLocal(local)
	conv := onb.ToWorld(world)
	if !local.Equals(conv) {
		t.Errorf("Invalid onb conversion \n src=%s \n conv=%s",
			local.ToString(), conv.ToString())
	}
}

func TestVector3_Dot(t *testing.T) {
	b := NewVector3(Sqrtf(3.0)/2.0, 0.5, 0.0).Normalize()
	a := NewVector3(0.5, Sqrtf(3.0)/2.0, 0.0).Normalize()

	if cos := a.Dot(b); !NearEqualEps(cos, Cosf(M_PIf/6.0), 0.01) {
		t.Errorf("Vector Dot is broken - %f != %f", cos, Cosf(M_PIf/6.0))
	}
}

func TestVector3_Arithmetic(t *testing.T) {
	v1 := randV3()
	v2 := randV3()

	v3 := v1.Mul(v2).Addf(1.0)

	asv := v3.Addf(-1.0).Mul(NewVector3(1.0/v2.x, 1.0/v2.y, 1.0/v2.z))

	if !NearEqualEps(v1.x, asv.x, 1e-5) {
		t.Errorf("Invalid arithmetic %f %f", v1.x, asv.x)
	}
}

func TestVector3_Len(t *testing.T) {
	v := randV3().Mulf(2.0)
	if v.Len() < 1.0 {
		t.Errorf("Invalid Len %f value (should be > 1)", v.Len())
	}
	v = v.Normalize()
	if v.Len() > 1.0002 {
		t.Errorf("Invalid Len %f value (should be < 1)", v.Len())
	}

}

func TestAABB_Contains(t *testing.T) {
	box := NewAABB(CreateVector3(-1.0), CreateVector3(1))

	for i := 0; i < 100; i++ {
		pt := randV3()
		if !box.Contains(pt) {
			t.Error("Contains broken")
		}
	}
}

func TestAABB_Union(t *testing.T) {
	box := NewAABB(CreateVector3(-1.0), CreateVector3(1))
	for i := 0; i < 100; i++ {
		pt := randV3().Mulf(float32(i + 1.0))
		if !box.Contains(pt) {
			box.Union(pt)
			box.Expand(1.0)
			if !box.Contains(pt) {
				t.Error("Union broken")
			}
		}
	}
}
