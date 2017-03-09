package corelib

import "math"

var M_PIf float32 = 3.14159265358979323846

var InvPIf float32 = 0.31830988618379067154

var Epsilon float32 = 1e-14

func Min(a, b float32) float32 {
	if a > b {
		return b
	}
	return a
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Clamp(a, min, max float32) float32 {
	return Min(max, Max(min, a))
}

func Sqrtf(value float32) float32 {
	u := math.Float32bits(value)
	i := (1 << 29) + (u >> 1) - (1 << 22) - 0x4C000
	return math.Float32frombits(i)
}

func NearEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) < float64(Epsilon)
}

func NearEqualEps(a, b, eps float32) bool {
	return math.Abs(float64(a-b)) < float64(eps)
}

func Sinf(x float32) (sin float32) {
	/*sin = 0.0
	if x < -3.14159265 {
		x += 6.28318531
	} else if x > 3.14159265 {
		x -= 6.28318531
	}
	if x < 0 {
		sin = 1.27323954*x + .405284735*x*x

		if sin < 0 {
			sin = .225*(sin*-sin-sin) + sin
		} else {
			sin = .225*(sin*sin-sin) + sin
		}
	} else {
		sin = 1.27323954*x - 0.405284735*x*x

		if sin < 0 {
			sin = .225*(sin*-sin-sin) + sin
		} else {
			sin = .225*(sin*sin-sin) + sin
		}
	}
	return sin*/
	sin = float32(math.Sin(float64(x)))
	return sin
}

func Cosf(x float32) (cos float32) {
	/*cos = 0.0
	if x < -3.14159265 {
		x += 6.28318531
	} else if x > 3.14159265 {
		x -= 6.28318531
	}

	if x < 0 {
		cos = 1.27323954*x + 0.405284735*x*x

		if cos < 0 {
			cos = .225*(cos*-cos-cos) + cos
		} else {
			cos = .225*(cos*cos-cos) + cos
		}
	} else {
		cos = 1.27323954*x - 0.405284735*x*x
		if cos < 0 {
			cos = .225*(cos*-cos-cos) + cos
		} else {
			cos = .225*(cos*cos-cos) + cos
		}
	}*/
	cos = float32(math.Cos(float64(x)))
	return cos
}

func Lerp(t, v1, v2 float32) float32 {
	return (1.0-t)*v1 + t*v2
}
