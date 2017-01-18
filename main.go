package main

import (
	"runtime"

	qp "github.com/Ariemeth/quantum-pulse"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {

	engine := new(qp.Engine)

	engine.Init()

	engine.LoadScene(engine.LoadSceneFile("scene1.json"))

	engine.Run()
}
