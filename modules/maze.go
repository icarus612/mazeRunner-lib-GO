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
		openPoints path = make(path)
		s          point
		e          point
		height     int    = build[0]
		width      int    = build[1]
		l          layout = m.layout
	)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := point{x, y}
			n := node{
				location: p,
			}
			if p[0] == 0 || p[1] == 0 || p[0] == height-1 || p[1] == width-1 {
				n.value = m.wallChar
			} else {
				openPoints[p] = true
				rng := rand.Int() % 3
				if rng%2 == 1 {
					n.value = m.wallChar
				} else {
					n.value = m.openChar
				}
			}
			l[x][y] = n
		}
	}
	switch buildType {
	case 'h':
		s = point{1, (rand.Int() % (width - 1)) + 1}
		e = point{height - 2, (rand.Int() % (width - 1)) + 1}
		delete(openPoints, s)
		delete(openPoints, e)
	case 'v':
		s = point{(rand.Int() % (height - 1)) + 1, 1}
		e = point{(rand.Int() % (height - 1)) + 1, width - 2}
		delete(openPoints, s)
		delete(openPoints, e)
	default:
		s = openPoints.toSlice()[rand.Int()%len(openPoints)]
		delete(openPoints, s)
		e = openPoints.toSlice()[rand.Int()%len(openPoints)]
		delete(openPoints, e)
	}
	l[s[0]][s[1]].value = m.startChar
	l[e[0]][e[1]].value = m.endChar
}

func (m *maze) ViewLayout() {
	for _, x := range m.layout {
		for _, y := range x {
			fmt.Print(string(y.value))
		}
		fmt.Println()
	}
}

func Maze(build [2]int, buildType rune) maze {
	//if len(build) == 2 {
	//	return maze[floor]
	//} else {
	//	return maze[dungon]
	//}

	l := make(layout, build[1])
	for i := range l {
		l[i] = make([]node, build[0])
	}

	m := maze{
		startChar: 's',
		endChar:   'e',
		wallChar:  '#',
		openChar:  ' ',
		layout:    l,
	}
	m.BuildNew(build, buildType)
	return m
}
