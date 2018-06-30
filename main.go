package main

import (
	"fmt"

	qpe "github.com/Ariemeth/quantum-pulse/engine"
)

const (
	screenWidth  = 800
	screenHeight = 600
	windowTitle  = "frame assault 2"
)

func main() {

	engine := new(qpe.Engine)

	engine.Init(screenWidth, screenHeight, windowTitle)

	sceneID, err := engine.LoadSceneFile("scene1.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	engine.LoadScene(sceneID)

	engine.Run()
}
