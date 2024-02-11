package lib

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
		s               point
		e               point
		length               = build[0]
		width                = build[1]
		height               = build[2]
		l                    = m.layout
		openPoints      path = make(path)
		localOpenPoints path = make(path)
		floorPoints     path = make(path)
		rls                  = rand.Int()%(length-2) + 1
		rws                  = rand.Int()%(width-2) + 1
		rle                  = rand.Int()%(length-2) + 1
		rwe                  = rand.Int()%(width-2) + 1
		fa                   = (length*width - 1) / 100
	)

	if height == 0 {
		height = 1
	}
	for z := 0; z < height; z++ {
		for y := 0; y < width; y++ {
			for x := 0; x < length; x++ {
				p := point{z, y, x}
				n := node{
					location: p,
				}
				if p[2] == 0 || p[1] == 0 || p[2] == length-1 || p[1] == width-1 {
					n.value = m.wallChar
				} else {
					rng := rand.Int() % 3
					if rng%2 == 1 {
						n.value = m.wallChar
					} else {
						localOpenPoints[p] = true
						n.value = m.openChar
					}
				}
				l[z][y][x] = n
			}
		}
		for i := 0; i <= fa && z < height-1; i++ {
			f := localOpenPoints.toSlice()[rand.Int()%len(localOpenPoints)]
			floorPoints[f] = true
			f[0]++
			floorPoints[f] = true

		}
		for k, v := range localOpenPoints {
			openPoints[k] = v
		}
		localOpenPoints = make(path)
	}

	for k, v := range floorPoints {
		if v {
			l[k[0]][k[1]][k[2]].value = m.floorChar
		}
	}

	//switch buildType {
	//case 'b':

	//default:
	//	s = openPoints.toSlice()[rand.Int()%len(openPoints)]
	//	delete(openPoints, s)
	//	e = openPoints.toSlice()[rand.Int()%len(openPoints)]
	//	delete(openPoints, e)
	//}

	s = point{0, rws, rls}
	e = point{height - 1, rwe, rle}
	delete(openPoints, s)
	delete(openPoints, e)

	l[s[0]][s[1]][s[2]].value = m.startChar
	l[e[0]][e[1]][e[2]].value = m.endChar
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
	for i := range l {
		l[i] = make(floor, b[1])
		for j := range l[i] {
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
	m.BuildNew(b, buildType)
	return m
}
