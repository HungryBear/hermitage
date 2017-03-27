package scenelib

import "hermitage/corelib"

type BaseSceneObject struct{
	Id, MaterialId int
	Name string
	World2Local, Local2World corelib.Transform
}

type Geometry struct{
	BaseSceneObject
	Shape corelib.Shape
}

type SceneCamera struct{
	BaseSceneObject
	Camera
}