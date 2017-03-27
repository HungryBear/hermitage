package corelib

type Shape interface {
	Intersect(ray *Ray) (bool, *RayHit)
	EvaluateNormal(position *Vector3) *Vector3
	Area() float32
	Bounds() *AABB
}

type Sphere struct {
	Position Vector3
	Radius   float32
}

type Parallelogram struct {
	Anchor Vector3
	V0, V1 Vector3
	Plane  Vector4
}

// Parallelogram

func NewParallelogram(a *Vector3, v0 *Vector3, v1 *Vector3) *Parallelogram {
	normal := v0
	normal = normal.Cross(v1).Normalize()

	d := normal.Dot(a)
	return &Parallelogram{Anchor: *a, V0: *v0.Normalize(), V1: *v1.Normalize(), Plane: *CreateVector4(normal, d)}
}

func (p *Parallelogram) EvaluateNormal(pos *Vector3) *Vector3 {
	return NewVector3(p.Plane.X, p.Plane.Y, p.Plane.Z)
}

func (p *Parallelogram) Bounds() *AABB{
	return NewAABB(NewVector3(Min(p.Anchor.X, Min(p.V0.X, p.V1.X)),Min(p.Anchor.Y, Min(p.V0.Y, p.V1.Y)),Min(p.Anchor.Z, Min(p.V0.Z, p.V1.Z))),
		NewVector3(Max(p.Anchor.X, Max(p.V0.X, p.V1.X)),Max(p.Anchor.Y, Max(p.V0.Y, p.V1.Y)),Max(p.Anchor.Z, Max(p.V0.Z, p.V1.Z))))
}

func (pl *Parallelogram) Intersect(ray *Ray) (bool, RayHit) {
	n := NewVector3(pl.Plane.X, pl.Plane.Y, pl.Plane.Z)
	dt := ray.Direction.Dot(n)
	t := (pl.Plane.W - n.Dot(&ray.Origin)) / dt
	hit := NewRayHit(0, 0, 1e10, InvalidHit)

	if t < ray.TMax && t > ray.TMin {
		p := ray.Origin.Add(ray.Direction.Copy().Mulf(t))
		vi := p.Sub(&pl.Anchor)
		a1 := pl.V0.Dot(vi)
		if a1 >= 0 && a1 <= 1 {
			a2 := pl.V1.Dot(vi)
			if a2 >= 0 && a2 <= 1 {
				hit.Distance = t
				return true, *hit
			}
		}

	}
	return false, *hit
}

func (p *Parallelogram) Area() float32 {
	tv1 := p.V0.Mulf(1.0 / p.V0.Dot(&p.V0))
	tv2 := p.V1.Mulf(1.0 / p.V1.Dot(&p.V1))
	return tv1.Cross(tv2).Len()
}

// Sphere

func NewSphere(pos *Vector3, rad float32) *Sphere {
	return &Sphere{Position: *pos, Radius: rad}
}

func (s*Sphere)Bounds()*AABB{
	return NewAABB(s.Position.Addf(-s.Radius), s.Position.Addf(s.Radius))
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

	hit := NewRayHit(0, 0, 1e10, InvalidHit)

	if discriminant > 0 {
		temp := (-b - Sqrtf(discriminant)) / a
		if temp < ray.TMax && temp > ray.TMin {
			hit.Distance = temp
			return true, *hit
		}
		temp = (-b + Sqrtf(discriminant)) / a
		if temp < ray.TMax && temp > ray.TMin {
			hit.Distance = temp
			return true, *hit
		}
	}
	return false, *hit
}

func (s *Sphere) Area() float32 {
	return 4.0 * M_PIf * s.Radius * s.Radius
}
