package coreimg

import (
	"math"
	"github.com/hungrybear/hermitage/corelib"
)

const (
	SPD_FLAGS_NONE = 0 << iota
	SPD_FLAGS_REFLECTION
	SPD_FLAGS_LIGHT
)

type SpectralDistribution interface {
	Sample(lambda float32) float32
}

type Spd struct{
	Samples 		[]float32
	Wavelength 		[]float32
	Flags			int
	LambdaMin, LambdaMax 	float32
	delta, inv_delta 	float32

}

func (s*Spd) Sample(lambda float32) float32{
	if lambda > s.LambdaMax || lambda < s.LambdaMin {
		return 0.0
	}

	x:=(lambda - s.LambdaMin) * s.delta;

	b0:=math.Floor(float64(x))
	b1:=math.Min(b0 + 1 , float64(len(s.Samples)-1))
	dx:= x - float32(b0)
	return corelib.Lerp(dx, s.Samples[int(b0)], s.Samples[int(b1)])
}

func (s*Spd) ToXyz() *RgbSpectrum{
	var x,y,z float32
	for i:=0;i<len(s.Samples);i++  {
		v:=s.Sample(s.Wavelength[i])
		x+=XFit_1931(s.Wavelength[i])*v
		y+=YFit_1931(s.Wavelength[i])*v
		z+=ZFit_1931(s.Wavelength[i])*v
	}
	return NewRgbSpectrum(x*683.0, y*683.0, z*683.0, RgbFlagsXyz)
}

func (s*Spd) Whitepoint(temp float32){
	bbvals := make([]float64, len(s.Samples))
	w:=1e-9*float64(s.LambdaMin)
	for i := 0; i<len(s.Samples);i++  {
		bbvals[i]=(4e-9*(3.74183e-16*math.Pow(w, -5.0))) / (math.Exp(1.4388e-2 / (w * float64(temp))) - 1.0)
		w+=1e-9*float64(s.delta)
	}

	var max float64
	for i := 0; i<len(s.Samples);i++  {
		if bbvals[i] > max{
			max=bbvals[i]
		}
	}

	for i := 0; i<len(s.Samples);i++ {
		s.Samples[i]=float32((1.0 / max)*bbvals[i])
	}
}

type RegularSpd struct{
	Spd
}

type IrregularSpd struct{
	Spd
}

func expf(e float32) float32{
	return float32(math.Exp(float64(e)))
}

func XFit_1931(wave float32)float32{
	var t1, t2, t3 float32
	if t1=(wave - 442.0);wave < 442{
		t1*=0.0624
	}else{
		t1*=0.0374
	}

	if t2=wave - 599.8;wave<599.8{
		t2*=0.0264
	}else{
		t2*=0.0323
	}

	if t3=wave - 501.1;wave<501.1{
		t3*=0.0490
	}else{
		t3*=0.0382
	}
	return 0.362*expf(-0.5*t1*t1) + 1.056*expf(-0.5*t2*t2)- 0.065*expf(-0.5*t3*t3);
}

func YFit_1931(wave float32)float32{
	var t1, t2 float32
	if t1=wave - 568.8;t1<568.8{
		t1*=0.0213
	}else{
		t1*=0.0247
	}

	if t2=wave - 530.9;wave<530.9{
		t2*=0.0613
	}else{
		t2*=0.0322
	}
	return 0.821*expf(-0.5*t1*t1) + 0.286*expf(-0.5*t2*t2);
}

func ZFit_1931(wave float32) float32 {
	var t1, t2 float32
	if t1=wave - 437.0;wave < 437.0{
		t1*=0.0845
	}else{
		t1*=0.0278
	}
	if t2=wave - 459.0;wave < 459.0{
		t2*=0.0385
	}else{
		t2*=0.0725
	}
	return 1.217*expf(-0.5*t1*t1) + 0.681*expf(-0.5*t2*t2);
}

func NewRegulardSpd(lambda, amplitude []float32, lambdaMin, lambdaMax float32, flags int) *RegularSpd{
	return &RegularSpd{
		Spd{
			LambdaMin:lambdaMin,
			LambdaMax:lambdaMax,
			Wavelength:lambda,
			Samples:amplitude,
			delta:(lambdaMax - lambdaMin) / float32(len(amplitude) - 1),
			inv_delta:1.0 /((lambdaMax - lambdaMin) / float32(len(amplitude) - 1)),
			Flags:flags}}
}


