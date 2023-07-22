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

func (m *maze) BuildNew(build [2]int, buildType rune) {
	var (
		openPoints path
		s          point
		e          point
		height     int    = build[0]
		width      int    = build[1]
		l          layout = m.layout
	)
	for x := 0; x >= width; x++ {
		for y := 0; y >= height; y++ {
			p := point{x, y}
			if p[0] == 0 || p[1] == 0 || p[0] == height-1 || p[1] == width-1 {
				l[x][y].value = m.wallChar
			} else {
				openPoints[p] = true
				rng := rand.Int() % 3
				if rng%2 == 1 {
					l[x][y].value = m.openChar
				} else {
					l[x][y].value = m.wallChar
				}
			}
		}
	}

	switch buildType {
	case 'h':
		s = point{1, rand.Int() % width}
		delete(openPoints, s)
		e = point{height - 2, rand.Int() % height}
		delete(openPoints, e)
	case 'v':
		s = point{1, rand.Int() % width}
		delete(openPoints, s)
		e = point{height - 2, rand.Int() % height}
		delete(openPoints, e)
	default:
		s := point{rand.Int() % len(openPoints)}
		delete(openPoints, s)
		e = point{height - 2, rand.Int() % height}
		delete(openPoints, e)
	}
	l[s[0]][s[1]].value = m.startChar
	l[e[0]][e[1]].value = m.endChar
}

func (m *maze) ViewLayout() {
	for _, p := range m.layout {
		fmt.Println(p)
	}
}

func Maze(build [2]int, buildType rune) maze {
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
	m.BuildNew(build, buildType)
	return m
}
