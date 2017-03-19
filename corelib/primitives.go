package corelib

import (
	"fmt"
	"math"
)

const InvalidHit uint32 = 0xffffff

type Vector2 struct {
	X, Y float32
}

type Vector3 struct {
	X, Y, Z float32
}

var X_Axis Vector3 = Vector3{1, 0, 0}
var Y_Axis Vector3 = Vector3{0, 1, 0}
var Z_Axis Vector3 = Vector3{0, 0, 1}

type Vector4 struct {
	X, Y, Z, W float32
}

type AABB struct {
	Min, Max Vector3
}

type Ray struct {
	Origin    Vector3
	Direction Vector3
	TMin      float32
	TMax      float32
	Time      float32
}

type RayHit struct {
	U, V     float32
	Distance float32
	Index    uint32
}

type Onb struct {
	Mx, My, Mz Vector3
}

//Axis Aligned Bounding Box

func NewAABB(min *Vector3, max *Vector3) *AABB {
	return &AABB{Min: *min, Max: *max}
}

func (b *AABB) Expand(f float32) *AABB {
	b.Min = *b.Min.Addf(-f)
	b.Max = *b.Max.Addf(f)
	return b
}

func (b *AABB) Overlap(b2 *AABB) bool {
	x := (b2.Max.X >= b.Min.X) && (b2.Min.X <= b.Max.X)
	y := (b2.Max.Y >= b.Min.Y) && (b2.Min.Y <= b.Max.Y)
	z := (b2.Max.Z >= b.Min.Z) && (b2.Min.Z <= b.Max.Z)
	return x || y || z
}

func (b *AABB) Union(p *Vector3) *AABB {
	if b.Min.X > p.X {
		b.Min.X = p.X
	}

	if b.Min.Y > p.Y {
		b.Min.Y = p.Y
	}

	if b.Min.Z > p.Z {
		b.Min.Z = p.Z
	}

	b.Max.X = Max(b.Max.X, p.X)
	b.Max.Y = Max(b.Max.Y, p.Y)
	b.Max.Z = Max(b.Max.Z, p.Z)

	return b
}

func (b *AABB) Contains(pt *Vector3) bool {
	return pt.X >= b.Min.X && pt.X <= b.Max.X &&
		pt.Y >= b.Min.Y && pt.Y <= b.Max.Y &&
		pt.Z >= b.Min.Z && pt.Z <= b.Max.Z
}

// Vector2

func NewVector2(x, y float32) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func CreateVector2(a float32) *Vector2 {
	return &Vector2{X: a, Y: a}
}

// Vector3

func Unit() *Vector3 {
	return &Vector3{1, 1, 1}
}

func Zero() *Vector3 {
	return &Vector3{0, 0, 0}
}

func CreateVector3(a float32) *Vector3 {
	return &Vector3{X: a, Y: a, Z: a}
}

func (v *Vector3) Copy() *Vector3 {
	return &Vector3{X: v.X, Y: v.Y, Z: v.Z}
}

func NewVector3(x, y, z float32) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}

func (v *Vector3) Add(v2 *Vector3) *Vector3 {
	return &Vector3{X: v.X + v2.X, Y: v.Y + v2.Y, Z: v.Z + v2.Z}
}

func (v *Vector3) Addf(v2 float32) *Vector3 {
	return &Vector3{X: v.X + v2, Y: v.Y + v2, Z: v.Z + v2}
}

func (v *Vector3) Sub(v2 *Vector3) *Vector3 {
	return &Vector3{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

func (v *Vector3) Mul(v2 *Vector3) *Vector3 {
	return &Vector3{X: v.X * v2.X, Y: v.Y * v2.Y, Z: v.Z * v2.Z}
}

func (v *Vector3) Mulf(v2 float32) *Vector3 {
	return &Vector3{X: v.X * v2, Y: v.Y * v2, Z: v.Z * v2}
}

func (v *Vector3) Dot(v2 *Vector3) float32 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (u *Vector3) Cross(v *Vector3) *Vector3 {
	return &Vector3{X: u.Y*v.Z - u.Z*v.Y, Y: u.Z*v.X - u.X*v.Z, Z: u.X*v.Y - u.Y*v.X}
}

func (v *Vector3) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v *Vector3) Normalize() *Vector3 {
	return v.Mulf(1.0 / v.Len())
}

func (v *Vector3) ToString() string {
	return fmt.Sprintf("%f %f %f ", v.X, v.Y, v.Z)
}

func (v *Vector3) Equals(v2 *Vector3) bool {
	return NearEqualEps(v.X, v2.X, 1e-4) && NearEqualEps(v.Y, v2.Y,1e-4) && NearEqualEps(v.Z, v2.Z,1e-4)
}

// Vector 4 methods

func NewVector4(x, y, z, w float32) *Vector4 {
	return &Vector4{X: x, Y: y, Z: z, W: w}
}

func CreateVector4(v *Vector3, f float32) *Vector4 {
	return &Vector4{X: v.X, Y: v.Y, Z: v.Z, W: f}
}

func (v *Vector4) Add(v2 *Vector4) *Vector4 {
	return &Vector4{X: v.X + v2.X, Y: v.Y + v2.Y, Z: v.Z + v2.Z, W: v.W + v2.W}
}

func (v *Vector4) Mul(v2 *Vector4) *Vector4 {
	return &Vector4{X: v.X * v2.X, Y: v.Y * v2.Y, Z: v.Z * v2.Z, W: v.W * v2.W}
}

func (v *Vector4) Mulf(v2 float32) *Vector4 {
	return &Vector4{X: v.X * v2, Y: v.Y * v2, Z: v.Z * v2, W: v.W * v2}
}

// Ray methods

func NewRay(pos, dir *Vector3) *Ray {
	return &Ray{Origin: *pos, Direction: *dir, TMin: 1e-4, TMax: 1e10}
}

func (ray *Ray) Point(dist float32) *Vector3 {
	return ray.Origin.Add(ray.Direction.Mulf(dist))
}

// RayHit

func NewRayHit(a, b, t float32, i uint32) *RayHit {
	return &RayHit{U: a, V: b, Distance: t, Index: i}
}

func Miss(hit *RayHit) bool {
	return hit.Index == InvalidHit
}

// Onb

func NewOnb(z Vector3) *Onb {
	tz := z.Normalize()
	var x Vector3
	if Abs(tz.X) > 0.99 {
		x = *Y_Axis.Copy()
	} else {
		x = *X_Axis.Copy()
	}
	ty := tz.Cross(&x).Normalize()
	tx := ty.Cross(tz)
	return &Onb{Mx: *tx, My: *ty, Mz: *tz}
}

func (o *Onb) ToWorld(v *Vector3) *Vector3 {
	return o.Mx.Mulf(v.X).Add(o.My.Mulf(v.Y)).Add(o.Mz.Mulf(v.Z))
}

func (o *Onb) ToLocal(v *Vector3) *Vector3 {
	return NewVector3(v.Dot(&o.Mx), v.Dot(&o.My), v.Dot(&o.Mz))
}
