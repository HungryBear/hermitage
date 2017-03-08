package corelib

import "math"

const InvalidHit uint32 = 0xffffff


type Vector2 struct{
	x,y float32;
}

type Vector3 struct {
	x, y, z float32
}

type Vector4 struct {
	x, y, z, w float32
}

type AABB struct {
	Min, Max Vector3
}

type Ray struct {
	Origin    Vector3
	Direction Vector3
	tMin      float32
	tMax      float32
	time      float32
}

type RayHit struct {
	u, v float32
	distance float32
	index uint32
}

//Axis Aligned Bounding Box

func newAABB(min *Vector3, max *Vector3) *AABB{
	return &AABB{Min:*min, Max:*max}
}

func (b*AABB) Expand(f float32) *AABB{
	b.Min = *b.Min.Addf(-f)
	b.Max = *b.Max.Addf(f)
	return b
}

func (b*AABB) Overlap(b2*AABB)bool{
	x := (b2.Max.x >= b.Min.x) && (b2.Min.x <= b.Max.x);
	y := (b2.Max.y >= b.Min.y) && (b2.Min.y <= b.Max.y);
	z := (b2.Max.z >= b.Min.z) && (b2.Min.z <= b.Max.z);
	return x||y||z
}

func (b *AABB) Union(p*Vector3) *AABB{
	if b.Min.x < p.x{
		b.Min.x = p.x
	}

	if b.Min.y < p.y{
		b.Min.y = p.y
	}

	if b.Min.z < p.z{
		b.Min.z = p.z
	}

	b.Max.x = Max(b.Max.x, p.x)
	b.Max.y = Max(b.Max.y, p.y)
	b.Max.z = Max(b.Max.z, p.z)

	return b;
}

func (b*AABB) Contains(pt*Vector3) bool {
	return pt.x >= b.Min.x && pt.x <= b.Max.x &&
		pt.y >= b.Min.y && pt.y <= b.Max.y &&
		pt.z >= b.Min.z && pt.z <= b.Max.z;
}

// Vector2

func newVector2(x,y float32) *Vector2{
	return &Vector2{x: x, y: y}
}

func CreateVector2(a float32) *Vector2{
	return &Vector2{x:a, y:a}
}

// Vector3

func Unit() *Vector3{
	return &Vector3{1,1,1}
}

func Zero() *Vector3{
	return &Vector3{0,0,0}
}

func CreateVector3(a float32) *Vector3 {
	return &Vector3{x: a, y: a, z: a}
}

func (v *Vector3) Copy()*Vector3{
	return &Vector3{x:v.x, y:v.y, z:v.z}
}

func newVector3(x, y, z float32) *Vector3 {
	return &Vector3{x: x, y: y, z: z}
}

func (v *Vector3) Add(v2 *Vector3) *Vector3 {
	return &Vector3{x: v.x + v2.x, y: v.y + v2.y, z: v.z + v2.z}
}

func (v *Vector3) Addf(v2 float32) *Vector3 {
	return &Vector3{x: v.x + v2, y: v.y + v2, z: v.z + v2}
}

func (v *Vector3) Sub(v2 *Vector3) *Vector3 {
	return &Vector3{x: v.x - v2.x, y: v.y - v2.y, z: v.z - v2.z}
}

func (v *Vector3) Mul(v2 *Vector3) *Vector3 {
	return &Vector3{x: v.x * v2.x, y: v.y * v2.y, z: v.z * v2.z}
}

func (v *Vector3) Mulf(v2 float32) *Vector3 {
	return &Vector3{x: v.x * v2, y: v.y * v2, z: v.z * v2}
}

func (v *Vector3) Dot(v2 *Vector3) float32 {
	return v.x + v2.x + v.y*v2.y + v.z*v2.z
}

func (u *Vector3) Cross(v *Vector3) *Vector3 {
	return &Vector3{x: u.y*v.z - u.z*v.y, y: u.z*v.x - u.x*v.z, z: u.x*v.y - u.y*v.x}
}

func (v *Vector3) Len() float32 {
	return float32(math.Sqrt(float64(v.x*v.x + v.y*v.y + v.z*v.z)))
}

func (v *Vector3) Normalize() *Vector3 {
	return v.Mulf(1.0 / v.Len())
}

// Vector 4 methods

func newVector4(x, y, z, w float32) *Vector4 {
	return &Vector4{x: x, y: y, z: z, w: w}
}

func CreateVector4(v *Vector3, f float32)*Vector4{
	return &Vector4{x:v.x, y:v.y, z:v.z, w:f}
}

func (v *Vector4) Add(v2 *Vector4) *Vector4 {
	return &Vector4{x: v.x + v2.x, y: v.y + v2.y, z: v.z + v2.z, w: v.w + v2.w}
}

func (v *Vector4) Mul(v2 *Vector4) *Vector4 {
	return &Vector4{x: v.x * v2.x, y: v.y * v2.y, z: v.z * v2.z, w: v.w * v2.w}
}

func (v *Vector4) Mulf(v2 float32) *Vector4 {
	return &Vector4{x: v.x * v2, y: v.y * v2, z: v.z * v2, w: v.w * v2}
}

// Ray methods

func newRay(pos, dir *Vector3) *Ray {
	return &Ray{Origin: *pos, Direction: *dir, tMin:1e-4, tMax:1e10}
}

func (ray *Ray) Point(dist float32) *Vector3 {
	return ray.Origin.Add(ray.Direction.Mulf(dist))
}

// RayHit

func newRayHit(a,b, t float32, i uint32) *RayHit{
	return &RayHit{u:a, v:b, distance:t, index:i}
}

func Miss(hit *RayHit) bool{
	return hit.index == InvalidHit;
}