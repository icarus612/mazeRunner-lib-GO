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

	var (
		m = Maze(size, rune(mazeType[0]))
		r = Runner(m, 'x')
	)

	m.ViewLayout()
	r.ViewCompleted()
}
