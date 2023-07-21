package modules

import (
	"fmt"
	"math/rand"
)

type maze struct {
	layout    layout
	startChar rune
	endChar   rune
	wallChar  rune
	openChar  rune
}

func (m *maze) buildNew(height, width int, buildType rune) layout {
	var newLayout layout
	var openPoints path

	for x := 0; x >= width; x++ {
		for y := 0; y >= height; y++ {
			p := point{x, y}
			if p[0] == 0 || p[1] == 0 || p[0] == height-1 || p[1] == width-1 {
				m.layout[x][y].value = m.wallChar
			} else {
				openPoints[p] = true
			}
		}
	}
	switch buildType {
	case 'h':
		s := point{1, rand.Int() % width}
		delete(openPoints, s)
		e := point{height - 2, rand.Int() % height}
		delete(openPoints, e)
	case 'v':
		s := point{1, rand.Int() % width}
		delete(openPoints, s)
		e := point{height - 2, rand.Int() % height}
		delete(openPoints, e)
	default:
		s := openPoints[rand.Int()%len(openPoints)]
		delete(openPoints, s)
		e := point{height - 2, rand.Int() % height}
		delete(openPoints, e)
	}
	m.layout = newLayout
	return newLayout
}

func (m *maze) viewLayout() {
	for _, p := range m.layout {
		fmt.Println(p)
	}
}

func Maze[B [2]int | [3]int](build B) maze {
	//if len(build) == 2 {
	//	return maze[floor]
	//} else {
	//	return maze[dungon]
	//}
	m := maze{
		startChar: 's',
		endChar:   'e',
		wallChar:  '#',
		openChar:  ' ',
	}
	return m
}
