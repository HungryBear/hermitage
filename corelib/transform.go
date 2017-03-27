package corelib

type Transform struct{
	m, mInv Matrix4x4
}

func NewTransform(mat *Matrix4x4)*Transform{
	return &Transform{m:*mat, mInv:*mat.Inverse()}
}

func (tr*Transform) TransformPoint(n *Vector3)*Vector3{
	xp:=tr.m[0][0] * n.X + tr.m[0][1] * n.Y + tr.m[0][2] * n.Z +tr.m[0][3]
	yp:=tr.m[1][0] * n.X + tr.m[1][1] * n.Y + tr.m[1][2] * n.Z +tr.m[1][3]
	zp:=tr.m[2][0] * n.X + tr.m[2][1] * n.Y + tr.m[2][2] * n.Z +tr.m[2][3]
	wp:=tr.m[3][0] * n.X + tr.m[3][1] * n.Y + tr.m[3][2] * n.Z +tr.m[3][3]

	return NewVector3(xp / wp, yp /wp, zp/wp)
}

func (tr*Transform) TransformDir(n*Vector3)*Vector3{
	xp:=tr.m[0][0] * n.X + tr.m[0][1] * n.Y + tr.m[0][2] * n.Z
	yp:=tr.m[1][0] * n.X + tr.m[1][1] * n.Y + tr.m[1][2] * n.Z
	zp:=tr.m[2][0] * n.X + tr.m[2][1] * n.Y + tr.m[2][2] * n.Z
	return NewVector3(xp, yp, zp)
}

func (tr*Transform) TransformNormal(n*Vector3)*Vector3{
	xp:=tr.mInv[0][0] * n.X + tr.mInv[0][1] * n.Y + tr.mInv[0][2] * n.Z
	yp:=tr.mInv[1][0] * n.X + tr.mInv[1][1] * n.Y + tr.mInv[1][2] * n.Z
	zp:=tr.mInv[2][0] * n.X + tr.mInv[2][1] * n.Y + tr.mInv[2][2] * n.Z
	return NewVector3(xp,yp,zp)
}

func (tr*Transform) Inverse()*Transform{
	return &Transform{m:tr.mInv, mInv:tr.m}
}

func CreateLookAt(pos, target, up *Vector3)*Transform{
	return NewTransform(LookAt(pos, target, up).Inverse())
}