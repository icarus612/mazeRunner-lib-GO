package modules

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type runner struct {
	completed    bool
	pathChar     rune
	openNodes    []rNode
	shortestPath path
	visited      []rNode
	toVisit      []rNode
	start        rNode
	end          rNode
	maze         maze
	mappedLayout layout
}

func (r *runner) getOpenNodes() {
	p := r.maze.layout
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[x]); y++ {
			newNode := p[x][y]
			if newNode.value != r.maze.wallChar {
				r.openNodes = append(r.openNodes, rNode{node: newNode})
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
				r.start = rNode{node: y}
			case e:
				r.end = rNode{node: y}
			}
		}
	}
}

func (r *runner) lookAround(n rNode) {
	nl := n.location
	for _, v := range r.openNodes {
		var (
			vl = v.location
		)
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
		rtv  = r.toVisit
		vtd  = r.visited
		vtdn = []point{}
	)
	for _, x := range vtd {
		vtdn = append(vtdn, x.location)
	}
	rtv = append(rtv, r.start)
	for len(rtv) > 0 {
		current := rtv[0]
		rtv = rtv[1:]

		if !slices.Contains(vtdn, current.location) {

			r.lookAround(current)
			vtd = append(vtd, current)
			for _, n := range current.children {
				if n.value == r.end.value {
					r.completed = true
				} else {
					rtv = append(rtv, n)
				}
			}
		}

	}
}

func (r *runner) buildPath() {
	var (
		start = r.start.location
		end   = r.end.location
		mpd   = r.mappedLayout
		m     = r.maze
		p     = r.pathChar
		s     = m.startChar
		e     = m.endChar
		w     = m.wallChar
		o     = m.openChar
		l     = m.layout
	)
	for slices.Contains([]rune{s, e, w, o}, p) {
		fmt.Println("The current path character can not be the same as the maze characters.")
		fmt.Printf("Current maze characters include %v, %v, %v, and %v.", s, e, w, o)
		fmt.Println("What would you like the new path the be?")
		fmt.Scan(&p)
	}

	mpd = l
	for _, x := range mpd {
		for _, y := range x {
			if start != y.location && end != y.location && slices.Contains(r.shortestPath.toSlice(), y.location) {
				y.value = p
			}
		}
	}
}

func (r *runner) setShortestPath(p path) {
	if len(p) < len(r.shortestPath) {
		r.shortestPath = p
	}
}

func (r runner) ViewCompleted() {
	for _, row := range r.mappedLayout {
		fmt.Println(row)
	}
}

func Runner(m maze, pathChar rune) runner {

	r := runner{
		completed:    false,
		maze:         m,
		pathChar:     pathChar | 'x',
		shortestPath: make(path),
	}
	r.getOpenNodes()
	r.findEndPoints()
	r.makeNodePaths()
	r.buildPath()
	return r
}
