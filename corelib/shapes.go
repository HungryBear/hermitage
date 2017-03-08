package corelib

type Shape interface {
	Intersect(ray *Ray) (bool, *RayHit)
	EvaluateNormal(position *Vector3) *Vector3
	Area() float32
}

type Sphere struct {
	Position Vector3
	Radius   float32
}

type Parallelogram struct {
	Anchor Vector3
	V0, V1 Vector3
	Plane Vector4
}


// Parallelogram

func newParallelogram(a * Vector3, v0 *Vector3 , v1 *Vector3) *Parallelogram{
	normal:= v0;
	normal = normal.Cross(v1).Normalize();

	d:=normal.Dot(a);
	return &Parallelogram{Anchor:*a, V0: *v0.Normalize(), V1:*v1.Normalize(), Plane:*CreateVector4(normal, d)}
}

func (p*Parallelogram) EvaluateNormal(pos *Vector3) *Vector3 {
	return newVector3(p.Plane.x, p.Plane.y, p.Plane.z)
}

func (pl*Parallelogram) Intersect(ray *Ray) (bool, RayHit) {
	n:= newVector3(pl.Plane.x, pl.Plane.y, pl.Plane.z)
	dt:=ray.Direction.Dot(n)
	t:=(pl.Plane.w - n.Dot(&ray.Origin)) / dt;
	hit := newRayHit(0, 0, 1e10, InvalidHit)

	if t < ray.tMax && t > ray.tMin {
		p:=ray.Origin.Add(ray.Direction.Copy().Mulf(t))
		vi:=p.Sub(&pl.Anchor)
		a1:=pl.V0.Dot(vi)
		if a1 >= 0 && a1 <= 1 {
			a2:=pl.V1.Dot(vi)
			if(a2 >= 0 && a2 <= 1) {
				hit.distance = t;
				return true, *hit
			}
		}

	}
	return false, *hit
}

func (p* Parallelogram) Area() float32{
	tv1  := p.V0.Mulf(1.0 / p.V0.Dot(&p.V0))
	tv2  := p.V1.Mulf(1.0 / p.V1.Dot(&p.V1))
	return tv1.Cross(tv2).Len()
}
// Sphere

func newSphere(pos *Vector3, rad float32) *Sphere {
	return &Sphere{Position: *pos, Radius: rad}
}

func (sphere *Sphere) EvaluateNormal(pos *Vector3) *Vector3 {
	return pos.Sub(&sphere.Position).Mulf(1.0 / sphere.Radius)
}

func (s *Sphere) Intersect(ray *Ray) (bool, RayHit) {
	oc := ray.Origin.Sub(&s.Position)
	a := ray.Direction.Dot(&ray.Direction)
	b := oc.Dot(&ray.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	hit := newRayHit(0, 0, 1e10, InvalidHit)

	if discriminant > 0 {
		temp := (-b - Sqrtf(discriminant)) / a
		if temp < ray.tMax && temp > ray.tMin {
			hit.distance = temp
			return true, *hit
		}
		temp = (-b + Sqrtf(discriminant)) / a
		if temp < ray.tMax && temp > ray.tMin {
			hit.distance = temp
			return true, *hit
		}
	}
	return false, *hit
}

func (s *Sphere) Area() float32 {
	return 4.0 * M_PIf * s.Radius * s.Radius
}
