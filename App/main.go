package main

import (
	"pathtracer"
)

func main(){
	pathtracer.initRenderer()
	pathtracer.render()
	pathtracer.exitRenderer()
}
