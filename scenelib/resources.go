package scenelib

import (
	"hermitage/coreimg"
)

type SceneResource interface {
	GetId() int
	Dispose() error
}

type SceneTexture struct{
	id int
	Image coreimg.Bitmap
}

type SpdData struct{
	id int

}
