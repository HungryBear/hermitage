package corelib

var primes = [...]int { 2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,
			83,89,97,101,103,107,109,113,127,131,137,139,149,151,157,163,167,173,179,181,
			191,193,197,199,211,223,227,229,233,239,241,251,257,263,269,271,277,281,283}

type HaltonSequence struct{
	Base, Seed uint32
}

func newHaltonSequence(base, seed int) *HaltonSequence{
	return &HaltonSequence{Base:uint32(base), Seed:uint32(seed)}
}

func (h*HaltonSequence) Initialize(seed uint32){
	h.Seed = seed >> 16;
}

func (h*HaltonSequence) NextFloat(index...int) float32{
	if h.Base >=60 {
		h.Base = 0;
		h.Seed++;
	}
	h.Base++;
	h.Seed++;
	return hal(int(h.Base-1), int(h.Seed-1))
}

func rev(i, p int) int{
	if i==0 {
		return i;
	}
	return p-i;
}

func hal(b, j int) float32{
	p:=primes[b]
	var h float32 =0.0
	f:=1.0 / float32(p)
	fct:=f
	for j > 0 {
		h += float32(rev(j%p, p))*fct
		j /=p
		fct*=f
	}
	return h
}