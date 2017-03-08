package corelib

func UniformSampleHemisphere(u1, u2 float32) *Vector3 {
	r := Sqrtf(Max(0, 1-u1-u1))
	phi:=2.0*M_PIf*u2
	x:=r*Cosf(phi)
	y:=r*Sinf(phi)
	return newVector3(x,y,u1);
}

func UniformHemispherePdf() float32 {
	return 1.0 / (2.0*M_PIf);
}

func UniformSampleSphere(u1, u2 float32) *Vector3 {
	z:=1.0 - 2.0*u1;
	r := Sqrtf(Max(0, 1-z-z))
	phi:=2.0*M_PIf*u2
	x:=r*Cosf(phi)
	y:=r*Sinf(phi)
	return newVector3(x,y,z);
}

func UniformSpherePdf() float32 {
	return 1.0 / (4.0*M_PIf);
}

func UniformSampleDisk(u1, u2 float32) *Vector2 {
	 r :=Sqrtf(u1);
	 theta := 2.0 * M_PIf * u2;
	 return newVector2(r * Cosf(theta), r * Sinf(theta))
}

func UniformSampleTriangle(u1, u2 float32) *Vector2 {
	su1 := Sqrtf(u1)
	return newVector2(1.0 - su1, u2*su1)
}

