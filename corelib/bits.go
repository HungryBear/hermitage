package corelib

func ExpandBits(v uint) uint {
	v = (v * 0x00010001) & 0xFF0000FF
	v = (v * 0x00000101) & 0x0F00F00F
	v = (v * 0x00000011) & 0xC30C30C3
	v = (v * 0x00000005) & 0x49249249
	return v
}

func Morton3D(x, y, z float32) uint {
	x = Min(Max(x*1024.0, 0.0), 1023.0)
	y = Min(Max(y*1024.0, 0.0), 1023.0)
	z = Min(Max(z*1024.0, 0.0), 1023.0)
	xx := ExpandBits(uint(x))
	yy := ExpandBits(uint(y))
	zz := ExpandBits(uint(z))
	return (xx << 2) | (yy << 1) | zz
}

func HasFlag(value, flag int) bool {
	return (value & flag) == value
}
