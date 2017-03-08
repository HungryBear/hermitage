package scenelib

import (
	"github.com/hungrybear/hermitage/corelib"
	"github.com/hungrybear/hermitage/coreimg"
)

const (
	HitType_Environment = 1 << iota
	HitType_Surface
	HitType_Lightsource
	HitType_Volume
)

const(
	BsdfEvent_Reflection = 1
	BsdfEvent_Transmission = 2
	BsdfEvent_Absorbtion = 4
	BsdfEvent_Emission = 8

	BsdfEvent_Diffuse = 16
	BsdfEvent_Specular = 32
	BsdfEvent_Glossy = 64
)

const(
	VolumeEvent_Absorbtion = 1
	VolumeEvent_Scattering = 2
	VolumeEvent_Emission = 4
)


type DifferentialGeometry struct{
	Position, DpDu, DpDv 	corelib.Vector3
	Ng, Ns, DnDu, DnDv 	corelib.Vector3
	UV 			corelib.Vector2
}

type IntersectionInfo struct {
	HitType   		int
	Position, IncomingDir 	corelib.Vector3
	Ng, Ns			corelib.Vector3
	TexCoord 		corelib.Vector2
}

type CameraSample struct{
	ImageX, ImageY	float32
	LensU,	LensV	float32
	Time 		float32
	EyeRay		corelib.Ray
}

type LightSample struct {
	LightId			int
	Spectra 		coreimg.RgbSpectrum
	LightRay		corelib.Ray
	Pdf, Distance, CosAt	float32
}

type BsdfSample struct {
	Spectra 		coreimg.RgbSpectrum
	IncomingDir		corelib.Vector3
	Pdf, CosAt		float32
	BsdfEvent		int
}