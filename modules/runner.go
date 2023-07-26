package modules

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type runner struct {
	Completed    bool
	pathChar     rune
	openNodes    []rNode
	shortestPath path
	visited      []point
	toVisit      []rNode
	start        rNode
	end          rNode
	maze         maze
	mappedLayout layout
}

func (r *runner) getOpenNodes() {
	p := r.maze.layout
	p.traverse(func(n node) {
		var (
			l       = n.location
			l0      = l[0]
			l1      = l[1]
			newNode = p[l0][l1]
		)
		if newNode.value != r.maze.wallChar {
			r.openNodes = append(r.openNodes, runNode(newNode))
		}
	})

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
				r.start = runNode(y)
			case e:
				r.end = runNode(y)
			}
		}
	}
}

func (r *runner) lookAround(n *rNode) {
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
		rtv = r.toVisit
		vtd = r.visited
	)
	rtv = append(rtv, r.start)
	for len(rtv) > 0 {
		var (
			current = rtv[0]
			cp      = current.path
			cl      = current.location
		)

		rtv = rtv[1:]
		if !slices.Contains(vtd, cl) {
			r.lookAround(&current)

			newPath := make(path, len(cp))
			for k, v := range cp {
				newPath[k] = v
			}

			newPath.add(cl)
			vtd = append(vtd, cl)
			for _, n := range current.children {
				n.path = newPath
				if n.value == r.end.value {
					r.Completed = true
					r.setShortestPath(newPath)
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
	)
	for slices.Contains([]rune{s, e, w, o}, p) {
		fmt.Println("The current path character can not be the same as the maze characters.")
		fmt.Printf("Current maze characters include %v, %v, %v, and %v.", s, e, w, o)
		fmt.Println("What would you like the new path the be?")
		fmt.Scan(&p)
	}

	mpd.traverse(func(n node) {
		var (
			l  = n.location
			l0 = l[0]
			l1 = l[1]
		)
		if start != l && end != l && slices.Contains(r.shortestPath.toSlice(), l) {
			mpd[l0][l1].value = p
		}
	})
}

func (r *runner) setShortestPath(p path) {
	if len(p) < len(r.shortestPath) || len(r.shortestPath) == 0 {
		r.shortestPath = p
	}
}

func (r *runner) ViewCompletedPath() {
	fmt.Println(r.shortestPath.toSlice())
}

func (r runner) ViewCompleted() {
	r.mappedLayout.print()
}

func Runner(m maze, pathChar rune) runner {
	r := runner{
		Completed:    false,
		maze:         m,
		mappedLayout: m.layout.deepCopy(),
		pathChar:     pathChar,
		shortestPath: make(path),
	}

	r.getOpenNodes()
	r.findEndPoints()
	r.makeNodePaths()
	r.buildPath()
	return r
}
