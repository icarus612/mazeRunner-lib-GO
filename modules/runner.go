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

func (r *runner) buildPathTree() {
	var (
		m      = r.maze
		l      = m.layout
		sc     = m.startChar
		ec     = m.endChar
		fc     = m.floorChar
		wc     = m.wallChar
		oc     = m.openChar
		oNodes = r.openNodes
	)
	l.traverse(
		func(n node) {
			if n.value == wc {
				return
			}
			rn := runNode(n)
			switch n.value {
			case sc:
				r.start = rn
			case ec:
				r.end = rn
			case oc, fc:
				oNodes = append(oNodes, rn)
				fallthrough
			case oc:
				r.checkSpace(&rn)
			case fc:
				r.checkStairs(&rn)
			}
		},
	)
}

func (r *runner) checkStairs(n *rNode) {
	var (
		m   = r.maze
		l   = m.layout
		nl  = n.location
		nl0 = nl[0]
		nl1 = nl[1]
		nl2 = nl[2]
		p1  = l[nl0][nl1][nl2]
	)
}

func (r *runner) checkSpace(n *rNode) {

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
