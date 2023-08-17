package main

import (
	"fmt"
	"mazerunner/modules"
)

func getSize(str string, val int, size int) int {
	fmt.Print(str)
	fmt.Scanln(&val)
	if val <= 0 {
		val = size
	}
	return val
}

func main() {
	var (
		size     [3]int
		mazeType string
		pathType string
		Maze     = modules.Maze
		Runner   = modules.Runner
	)
	fmt.Println("What is the L x W x H size you would like for the maze?")
	fmt.Println("Default values are 40 x 20 x 3 (press enter)?")

	size[0] = getSize("L: ", size[0], 40)
	size[1] = getSize("W: ", size[1], 20)
	size[2] = getSize("H: ", size[2], 3)
	fmt.Print("Maze type  ? ")
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
	if r.Completed {
		r.ViewCompleted()
	} else {
		fmt.Println("Can Not Complete!")
	}
}
