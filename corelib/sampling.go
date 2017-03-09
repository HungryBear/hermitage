package corelib

type SamplesCollection struct {
	Samples1DCount int
	Samples1D      []float32
	Samples2DCount int
	Samples2D      []Vector2
}

type Sampler interface {
	Get1DSample() float32
	Get2DSample() *Vector2
	RequestSamples(count1d, count2d int) *SamplesCollection
}

type RandomProvider interface {
	Initialize(seed uint32)
	NextFloat(index ...int) float32
}

const REAL_UINT_INT float64 = 0.00000000045
const Y uint32 = 842502087
const Z uint32 = 3579807591
const W uint32 = 273326509

type MersenneTwister struct {
	x, y, z, w uint32
}

type LcgRandom struct {
	seed uint32
}

type UniformSampler struct {
	Width, Height int
	Seed          int
	Random        RandomProvider
}

// UniformSampler
func NewUniformSampler(width, height, seed int, rnd RandomProvider) *UniformSampler {
	return &UniformSampler{Width: width, Height: height, Seed: seed, Random: rnd}
}

func (u *UniformSampler) Get1DSample() float32 {
	return u.Random.NextFloat()
}

func (u *UniformSampler) Get2DSample() *Vector2 {
	return NewVector2(u.Random.NextFloat(), u.Random.NextFloat())
}

func (u *UniformSampler) RequestSamples(count1D, count2D int) *SamplesCollection {
	result := NewSamplesCollection(count1D, count2D)

	for i := 0; i < count1D; i++ {
		result.Samples1D[i] = u.Random.NextFloat(u.Seed)
	}

	for i := 0; i < count2D; i++ {
		result.Samples2D[i] = *NewVector2(u.Random.NextFloat(u.Seed), u.Random.NextFloat(u.Seed))
	}
	return result
}

//SamplesCollection

func NewSamplesCollection(max1D, max2D int) *SamplesCollection {
	return &SamplesCollection{Samples1D: make([]float32, max1D), Samples1DCount: max1D, Samples2D: make([]Vector2, max2D), Samples2DCount: max2D}
}

//Linear congr. generator

func NewLcgRandom(newSeed int) *LcgRandom {
	return &LcgRandom{seed: uint32(newSeed)}
}

func (l *LcgRandom) Initialize(seed uint32) {
	l.seed = seed
}

func lcg(prev *uint32) uint32 {
	*prev = (*prev*8121 + 28411) % 134456
	return *prev
}

func (l *LcgRandom) NextFloat(index ...int) float32 {
	return float32(lcg(&l.seed)) / float32(0x020000)
}

// Mersenne twister
func NewMersenneTwister(seed int) *MersenneTwister {
	return &MersenneTwister{x: uint32(seed), y: Y, z: Z, w: W}
}

func (m *MersenneTwister) Initialize(seed uint32) {
	m.x = seed
	m.y = Y
	m.z = Z
	m.w = W
}

func (m *MersenneTwister) NextFloat(index ...int) float32 {
	t := m.x ^ (m.x << 11)
	m.x = m.y
	m.y = m.z
	m.z = m.w
	m.w = (m.w ^ (m.w >> 19)) ^ (t ^ (t >> 8))

	return float32(REAL_UINT_INT) * float32(uint32(0x7FFFFFFF&m.w))
}
