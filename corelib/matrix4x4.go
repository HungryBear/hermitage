package corelib

import "math"

type Matrix4x4 [4][4]float32

func NewMatrix4x4(m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44 float32) *Matrix4x4 {
	return &Matrix4x4{
		{m11, m12, m13, m14},
		{m21, m22, m23, m24},
		{m31, m32, m33, m34},
		{m41, m42, m43, m44},
	}
}

func det2x2(a, b, c, d float32) float32 {
	return a*d - b*c
}

func det3x3(a1, a2, a3, b1, b2, b3, c1, c2, c3 float32) float32 {
	return a1*det2x2(b2, b3, c2, c3) - b1*det2x2(a2, a3, c2, c3) + c1*det2x2(a2, a3, b2, b3)
}

func Identity() *Matrix4x4 {
	return NewMatrix4x4(1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)
}

func (m *Matrix4x4) Det() float32 {
	var a1, a2, a3, a4, b1, b2, b3, b4, c1, c2, c3, c4, d1, d2, d3, d4 float32
	a1 = m[0][0]
	b1 = m[1][0]
	c1 = m[2][0]
	d1 = m[3][0]
	a2 = m[0][1]
	b2 = m[1][1]
	c2 = m[2][1]
	d2 = m[3][1]
	a3 = m[0][2]
	b3 = m[1][2]
	c3 = m[2][2]
	d3 = m[3][2]
	a4 = m[0][3]
	b4 = m[1][3]
	c4 = m[2][3]
	d4 = m[3][3]

	return a1*det3x3(b2, b3, b4, c2, c3, c4, d2, d3, d4) -
		b1*det3x3(a2, a3, a4, c2, c3, c4, d2, d3, d4) +
		c1*det3x3(a2, a3, a4, b2, b3, b4, d2, d3, d4) -
		d1*det3x3(a2, a3, a4, b2, b3, b4, c2, c3, c4)
}

func (m *Matrix4x4) Transpose() *Matrix4x4 {
	return &Matrix4x4{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

func (m1 *Matrix4x4) Add(m2 *Matrix4x4)*Matrix4x4{
	return &Matrix4x4{ { m1[0][0] + m2[0][0],m1[0][1] + m2[0][1],m1[0][2] + m2[0][2], m1[0][3] + m2[0][3]},
		{m1[1][0] + m2[1][0],m1[1][1] + m2[1][1],m1[1][2] + m2[1][2], m1[1][3] + m2[1][3]},
		{m1[2][0] + m2[2][0],m1[2][1] + m2[2][1],m1[2][2] + m2[2][2], m1[2][3] + m2[2][3]},
		{m1[3][0] + m2[3][0],m1[3][1] + m2[3][1],m1[3][2] + m2[3][2], m1[3][3] + m2[3][3]},
	}
}

func (m1 *Matrix4x4) Sub(m2 *Matrix4x4)*Matrix4x4{
	return &Matrix4x4{ { m1[0][0] - m2[0][0],m1[0][1] - m2[0][1],m1[0][2] - m2[0][2], m1[0][3] - m2[0][3]},
			   {m1[1][0] - m2[1][0],m1[1][1] - m2[1][1],m1[1][2] - m2[1][2], m1[1][3] - m2[1][3]},
			   {m1[2][0] - m2[2][0],m1[2][1] - m2[2][1],m1[2][2] - m2[2][2], m1[2][3] - m2[2][3]},
			   {m1[3][0] - m2[3][0],m1[3][1] - m2[3][1],m1[3][2] - m2[3][2], m1[3][3] - m2[3][3]},
	}
}

func (m*Matrix4x4) Inverse() * Matrix4x4{
	var a1, a2, a3, a4, b1, b2, b3, b4 float32
	var c1, c2, c3, c4, d1, d2, d3, d4 float32
	det:= m.Det()
	a1 = m[0][0]
	b1 = m[1][0]
	c1 = m[2][0]
	d1 = m[3][0]
	a2 = m[0][1]
	b2 = m[1][1]
	c2 = m[2][1]
	d2 = m[3][1]
	a3 = m[0][2]
	b3 = m[1][2]
	c3 = m[2][2]
	d3 = m[3][2]
	a4 = m[0][3]
	b4 = m[1][3]
	c4 = m[2][3]
	d4 = m[3][3]
	r :=&Matrix4x4{}
	r[0][0] = det3x3(b2, b3, b4, c2, c3, c4, d2, d3, d4) / det;
	r[0][1] = -det3x3(a2, a3, a4, c2, c3, c4, d2, d3, d4) / det;
	r[0][2] = det3x3(a2, a3, a4, b2, b3, b4, d2, d3, d4) / det;
	r[0][3] = -det3x3(a2, a3, a4, b2, b3, b4, c2, c3, c4) / det;
	r[1][0] = -det3x3(b1, b3, b4, c1, c3, c4, d1, d3, d4) / det;
	r[1][1] = det3x3(a1, a3, a4, c1, c3, c4, d1, d3, d4) / det;
	r[1][2] = -det3x3(a1, a3, a4, b1, b3, b4, d1, d3, d4) / det;
	r[1][3] = det3x3(a1, a3, a4, b1, b3, b4, c1, c3, c4) / det;
	r[2][0] = det3x3(b1, b2, b4, c1, c2, c4, d1, d2, d4) / det;
	r[2][1] = -det3x3(a1, a2, a4, c1, c2, c4, d1, d2, d4) / det;
	r[2][2] = det3x3(a1, a2, a4, b1, b2, b4, d1, d2, d4) / det;
	r[2][3] = -det3x3(a1, a2, a4, b1, b2, b4, c1, c2, c4) / det;
	r[3][0] = -det3x3(b1, b2, b3, c1, c2, c3, d1, d2, d3) / det;
	r[3][1] = det3x3(a1, a2, a3, c1, c2, c3, d1, d2, d3) / det;
	r[3][2] = -det3x3(a1, a2, a3, b1, b2, b3, d1, d2, d3) / det;
	r[3][3] = det3x3(a1, a2, a3, b1, b2, b3, c1, c2, c3) / det;

	return r
}

func (m1*Matrix4x4) Mul(m2 *Matrix4x4)*Matrix4x4{
	r:=&Matrix4x4{}

	r[0][0] = m1[0][0] * m2[0][0] + m1[0][1] * m2[1][0] + m1[0][2] * m2[2][0] + m1[0][3] * m2[3][0]
	r[0][1] = m1[0][0] * m2[0][1] + m1[0][1] * m2[1][1] + m1[0][2] * m2[2][1] + m1[0][3] * m2[3][1]
	r[0][2] = m1[0][0] * m2[0][2] + m1[0][1] * m2[1][2] + m1[0][2] * m2[2][2] + m1[0][3] * m2[3][2]
	r[0][3] = m1[0][0] * m2[0][3] + m1[0][1] * m2[1][3] + m1[0][2] * m2[2][3] + m1[0][3] * m2[3][3]
	r[1][0] = m1[1][0] * m2[0][0] + m1[1][1] * m2[1][0] + m1[1][2] * m2[2][0] + m1[1][3] * m2[3][0]
	r[1][1] = m1[1][0] * m2[0][1] + m1[1][1] * m2[1][1] + m1[1][2] * m2[2][1] + m1[1][3] * m2[3][1]
	r[1][2] = m1[1][0] * m2[0][2] + m1[1][1] * m2[1][2] + m1[1][2] * m2[2][2] + m1[1][3] * m2[3][2]
	r[1][3] = m1[1][0] * m2[0][3] + m1[1][1] * m2[1][3] + m1[1][2] * m2[2][3] + m1[1][3] * m2[3][3]
	r[2][0] = m1[2][0] * m2[0][0] + m1[2][1] * m2[1][0] + m1[2][2] * m2[2][0] + m1[2][3] * m2[3][0]
	r[2][1] = m1[2][0] * m2[0][1] + m1[2][1] * m2[1][1] + m1[2][2] * m2[2][1] + m1[2][3] * m2[3][1]
	r[2][2] = m1[2][0] * m2[0][2] + m1[2][1] * m2[1][2] + m1[2][2] * m2[2][2] + m1[2][3] * m2[3][2]
	r[2][3] = m1[2][0] * m2[0][3] + m1[2][1] * m2[1][3] + m1[2][2] * m2[2][3] + m1[2][3] * m2[3][3]
	r[3][0] = m1[3][0] * m2[0][0] + m1[3][1] * m2[1][0] + m1[3][2] * m2[2][0] + m1[3][3] * m2[3][0]
	r[3][1] = m1[3][0] * m2[0][1] + m1[3][1] * m2[1][1] + m1[3][2] * m2[2][1] + m1[3][3] * m2[3][1]
	r[3][2] = m1[3][0] * m2[0][2] + m1[3][1] * m2[1][2] + m1[3][2] * m2[2][2] + m1[3][3] * m2[3][2]
	r[3][3] = m1[3][0] * m2[0][3] + m1[3][1] * m2[1][3] + m1[3][2] * m2[2][3] + m1[3][3] * m2[3][3]

	return r;
}

func Translate(x,y,z float32)*Matrix4x4{
	m := Identity()
	m[3][0] = x
	m[3][1] = y
	m[3][2] = z
	return m
}

func Scale(x,y,z float32)*Matrix4x4{
	m := Identity()
	m[0][0] = x
	m[1][1] = y
	m[2][2] = z
	return m
}

func RotateX(angle float32) *Matrix4x4{
	m:=Identity()
	cosa:=Cosf(angle)
	sina:=Sinf(angle)
	m[1][1]=cosa
	m[1][2]=sina
	m[2][1]=-sina
	m[2][2]=cosa
	return m
}

func RotateY(angle float32) *Matrix4x4{
	m:=Identity()
	cosa:=Cosf(angle)
	sina:=Sinf(angle)
	m[0][0]=cosa
	m[0][2]=-sina
	m[2][0]=sina
	m[2][2]=cosa
	return m
}

func RotateZ(angle float32) *Matrix4x4{
	m:=Identity()
	cosa:=Cosf(angle)
	sina:=Sinf(angle)
	m[0][0]=cosa
	m[0][1]=sina
	m[1][0]=-sina
	m[1][1]=cosa
	return m
}


func Perspective(aFov, aNear, aFar float32) *Matrix4x4 {
	f := 1.0 / float32(math.Tan(float64(aFov)*math.Pi/360.0))
	d := float32(1.0 / (aNear - aFar))
	return NewMatrix4x4(
		f, 0, 0, 0,
		0, -f, 0, 0,
		0, 0, (aNear+aFar)*d, 2.0*d*aNear*aFar,
		0, 0, -1.0, 0)
}

func LookAt(pos, target, up *Vector3)*Matrix4x4{
	m:=&Matrix4x4{}
	m[0][3] = pos.X
	m[1][3] = pos.Y
	m[2][3] = pos.Z
	m[3][3] = 1.0
	dir:=target.Sub(pos).Normalize()
	right:=dir.Cross(up).Normalize()
	newUp:=right.Cross(dir).Normalize()
	m[0][0] = right.X
	m[1][0] = right.Y
	m[2][0] = right.Z
	m[3][0] = 0
	m[0][1] = newUp.X
	m[1][1] = newUp.Y
	m[2][1] = newUp.Z
	m[3][1] = 0
	m[0][2] = dir.X
	m[1][2] = dir.Y
	m[2][2] = dir.Z
	m[3][2] = 0
	return m
}


