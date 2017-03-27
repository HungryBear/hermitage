package scenelib

type Scene struct{
	Solids 		[]Geometry
	Lights		[]LightSource
	Resources 	[]SceneResource
	Cameras    	[]Camera
	Name		string
}

