package scenelib

import "hermitage/corelib"

type Camera interface{
	Setup(pos, dir, up corelib.Vector3, fov float32, ...float32)
	GetRay(sample *CameraSample) corelib.Ray
}

type PinholeCamera struct{
	Position, Target, Up corelib.Vector3
	Fov float32
}
