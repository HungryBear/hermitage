package corelib

import "testing"

func TestTransform_TransformPoint(t *testing.T) {
	m:=NewTransform(randMatrix())
	mInv := m.Inverse()
	local := randV3().Mulf(10.0)

	world := m.TransformPoint(local)
	conv := mInv.TransformPoint(world)
	if !local.Equals(conv) {
		t.Errorf("Invalid point conversion \n src=%s \n conv=%s",
			local.ToString(), conv.ToString())
	}


}

func TestTransform_TransformDir(t *testing.T) {
	m:=NewTransform(randMatrix())
	mInv := m.Inverse()
	local := randV3().Normalize()

	world := m.TransformDir(local).Normalize()
	conv := mInv.TransformDir(world).Normalize()
	if !local.Equals(conv) {
		t.Errorf("Invalid dir conversion \n src=%s \n conv=%s \n world=%s",
			local.ToString(), conv.ToString(), world.ToString())
	}


}

func TestTransform_TransformNormal(t *testing.T) {
	m:=NewTransform(randMatrix())
	mInv := m.Inverse()
	local := randV3().Normalize()

	world := m.TransformNormal(local).Normalize()
	conv := mInv.TransformNormal(world).Normalize()
	if !local.Equals(conv) {
		t.Errorf("Invalid normal conversion \n src=%s \n conv=%s \n world=%s",
			local.ToString(), conv.ToString(), world.ToString())
	}


}