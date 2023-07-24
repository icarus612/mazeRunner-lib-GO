package main

import (
	"fmt"
	"mazerunner/modules"
)

func getSize(str string, val int) int {
	fmt.Print(str)
	fmt.Scanln(&val)
	if val <= 0 {
		val = 10
	}
	return val
}

func main() {
	var (
		size     [2]int
		mazeType string
		pathType string
		Maze     = modules.Maze
		Runner   = modules.Runner
	)
	fmt.Println("What is the M x N size you would like for the maze?")
	fmt.Println("Default values are 10 x 10 (press enter)?")

	size[0] = getSize("M: ", size[1])
	size[1] = getSize("N: ", size[1])
	fmt.Print("Maze type horizontal (h), vertical (v) or random (r) (enter for random)? ")
	fmt.Scanln(&mazeType)
	if len(mazeType) == 0 {
		mazeType = "r"
	}
	fmt.Print("What character should the path of the completed maze be (enter for x)? ")
	fmt.Scanln(&pathType)
	if len(pathType) == 0 {
		pathType = "x"
	}
	var (
		m = Maze(size, rune(mazeType[0]))
		r = Runner(m, rune(pathType[0]))
	)

	m.ViewLayout()
	r.ViewCompletedPath()
	if r.Completed {
		r.ViewCompleted()
	} else {
		fmt.Println("Can Not Complete!")
	}
}
