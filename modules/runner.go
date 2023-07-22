package modules

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type runner struct {
	completed     bool
	pathChar      rune
	openNodes     []node
	possiblePaths []path
	visited       path
	toVisit       path
	start         node
	end           node
	maze          maze
	mappedMaze    layout
}

func (r *runner) getOpenNodes() {
	p := r.maze.layout
	for x := 0; x < len(p); x++ {
		for y := 0; x < len(p[x]); y++ {
			newNode := p[x][y]
			if newNode.value != r.maze.wallChar {
				r.openNodes = append(r.openNodes, newNode)
			}
		}
	}
}

func (r *runner) findEndPoints() {
	var (
		m = r.maze
		l = m.layout
		s = m.startChar
		e = m.endChar
	)

	for _, x := range l {
		for _, y := range x {
			switch y.value {
			case s:
				r.start = y
			case e:
				r.end = y
			}
		}
	}
}

func (r *runner) lookAround(n node) {
	nl := n.location
	for _, v := range r.openNodes {
		vl := v.location
		if vl[0]-1 == nl[0] && vl[1] == nl[1] {
			n.addChild(v)
		} else if vl[0]+1 == nl[0] && vl[1] == nl[1] {
			n.addChild(v)
		} else if vl[1]-1 == nl[1] && vl[0] == nl[0] {
			n.addChild(v)
		} else if vl[1]+1 == nl[1] && vl[0] == nl[0] {
			n.addChild(v)
		}
	}
}

func (r *runner) makeNodePaths() {
	var (
		rtv = r.toVisit
	)
	rtv.add(r.start)
	for len(rtv) > 0 {
		for _, x := range rtv.toSlice() {
			delete(rtv, x)
			if !r.visited[x] {
				r.lookAround(rtv)
				r.visited.add(rtv)
				for i := range r.children {
					if i.value == r.end.value {
						r.completed = true
					} else {
						r.toVisit.add(i)
					}
				}
			}
		}
	}
}

func (r *runner) buildPath() {
	var (
		m     = r.maze
		start = r.start.location
		end   = r.end.location
		p     = r.pathChar
		s     = m.startChar
		e     = m.endChar
		w     = m.wallChar
		o     = m.openChar
		l     = m.layout
		mpd   = m.mappedMaze
	)
	for slices.Contains([]rune{s, e, w, o}, p) {
		fmt.Println("The current path character can not be the same as the maze characters.")
		fmt.Printf("Current maze characters include %s, %s, %s, and %s.", s, e, w, o)
		fmt.Println("What would you like the new path the be?")
		fmt.Scan(&p)
	}

	mpd = m
	for _, x := range mpd {
		for _, y := range x {
			if start != y.location && r.path.Contains(y.location) {
				y.value = p
			}
		}
	}
}

func (r *runner) setShortestPath(p path) {
	if len(p) < len(r.path) {
		r.path = p
	}
}

func (r runner) ViewCompleted() {
	for _, row := range r.mappedMaze {
		fmt.Println(row)
	}
}

func Runner(m maze, pathChar rune) runner {

	r := runner{
		completed: false,
		maze:      m,
		pathChar:  pathChar || x,
	}
	r.getOpenNodes()
	r.findEndPoints()
	r.makeNodePaths()
	r.buildPath()
	return r
}
