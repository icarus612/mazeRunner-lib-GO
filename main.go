package main

import (
	"fmt"
	"mazerunner/modules"
)

func main() {
	var (
		size   [2]int
		Maze   = modules.Maze
		Runner = modules.Runner
	)

	fmt.Println("What is the N x M size you would like for the maze?")
	fmt.Print("N: ")
	fmt.Scan(&size[0])
	fmt.Print("W: ")
	fmt.Scan(&size[1])
	var (
		m = Maze(size, 'h')
		r = Runner(m, 'x')
	)
	m.ViewLayout()
	r.ViewCompleted()
}
