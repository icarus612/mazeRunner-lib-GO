package modules

import (
	"math/rand"
)

type maze struct {
	layout    layout
	startChar rune
	endChar   rune
	floorChar rune
	wallChar  rune
	openChar  rune
}

func (m *maze) BuildNew(build [3]int, buildType rune) {
	var (
		openPoints path = make(path)
		s          point
		e          point

		height = build[2]
		width  = build[1]
		length = build[0]
		l      = m.layout
		rws    = rand.Int() % (height - 1)
		rhs    = rand.Int() % (width - 1)
		rwe    = rand.Int() % (height - 1)
		rhe    = rand.Int() % (width - 1)
	)

	if length == 0 {
		length = 1
	}
	for z := 0; z < length; z++ {
		for y := 0; y < width; y++ {
			for x := 0; x < height; x++ {
				p := point{z, y, x}
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
				l[z][y][x] = n
			}
		}
	}

	switch buildType {
	case 'b':

	default:
		s = openPoints.toSlice()[rand.Int()%len(openPoints)]
		delete(openPoints, s)
		e = openPoints.toSlice()[rand.Int()%len(openPoints)]
		delete(openPoints, e)
	}

	s = point{0, rws, rhs}
	e = point{height - 1, rwe, rhe}
	delete(openPoints, s)
	delete(openPoints, e)

	l[s[0]][s[1]].value = m.startChar
	l[e[0]][e[1]].value = m.endChar
}

func (m maze) ViewLayout() {
	m.layout.print()
}

func Maze(b [3]int, buildType rune) maze {
	//if len(build) == 2 {
	//	return maze[floor]
	//} else {
	//	return maze[dungon]
	//}
	if b[2] == 0 {
		b[2] = 1
	}

	l := make(layout, b[2])
	for i, f := range l {
		l[i] = make(floor, b[1])
		for j := range f {
			l[i][j] = make([]node, b[0])
		}
	}

	m := maze{
		startChar: 's',
		endChar:   'e',
		floorChar: 'f',
		wallChar:  '#',
		openChar:  ' ',
		layout:    l,
	}
	m.BuildNew(build, buildType)
	return m
}
