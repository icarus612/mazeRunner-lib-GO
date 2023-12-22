package mazerunner

import (
	"fmt"
	"slices"
)

type runner struct {
	Completed    bool
	pathChar     rune
	shortestPath path
	visited      []point
	toVisit      []rNode
	start        rNode
	end          rNode
	maze         maze
	mappedLayout layout
}

func (r *runner) findEndpoints() {
	var (
		m  = r.maze
		sc = m.startChar
		ec = m.endChar
	)
	m.layout.traverse(
		func(n node) {
			rn := runNode(n)
			switch n.value {
			case sc:
				r.start = rn
			case ec:
				r.end = rn
			}
		},
	)
}

func (r *runner) lookAround(n *rNode) {
	var (
		m  = r.maze
		fc = m.floorChar
		oc = m.openChar
		sc = m.startChar
	)

	switch n.value {
	case oc, sc, fc:
		r.checkSpace(n)
		fallthrough
	case fc:
		r.checkStairs(n)
	}

}

func (r *runner) checkStairs(n *rNode) {
	var (
		m   = r.maze
		l   = m.layout
		nl  = n.location
		fc  = m.floorChar
		nl0 = nl[0]
		nl1 = nl[1]
		nl2 = nl[2]
	)

	if nl0 > 0 {
		pf := l[nl0-1][nl1][nl2]
		if pf.value == fc {
			n.addChild(runNode(pf))
		}
	}
	if nl0 < len(l)-1 {
		pb := l[nl0+1][nl1][nl2]
		if pb.value == fc {
			n.addChild(runNode(pb))
		}
	}
}

func (r *runner) checkSpace(n *rNode) {
	var (
		m   = r.maze
		oc  = m.openChar
		fc  = m.floorChar
		sc  = m.startChar
		ec  = m.endChar
		nl  = n.location
		nl0 = nl[0]
		nl1 = nl[1]
		nl2 = nl[2]
		cf  = m.layout[nl0]
		f1  = cf[nl1-1][nl2]
		f2  = cf[nl1+1][nl2]
		f3  = cf[nl1][nl2-1]
		f4  = cf[nl1][nl2+1]
	)

	for _, x := range []node{f1, f2, f3, f4} {
		switch x.value {
		case oc, fc, sc, ec:
			n.addChild(runNode(x))
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
		mpd = r.mappedLayout
		m   = r.maze
		p   = r.pathChar
		s   = m.startChar
		e   = m.endChar
		w   = m.wallChar
		o   = m.openChar
		f   = m.floorChar
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
			v  = n.value
			l0 = l[0]
			l1 = l[1]
			l2 = l[2]
		)
		if !slices.Contains([]rune{s, e, f}, v) && slices.Contains(r.shortestPath.toSlice(), l) {
			mpd[l0][l1][l2].value = p
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
	r.findEndpoints()
	r.makeNodePaths()
	r.buildPath()
	return r
}
