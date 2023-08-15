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
				r.checkSpace(&rn)
				fallthrough
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
	)

	if nl0 > 0 {
		pf := l[nl0-1][nl1][nl2]
		if pf.value == m.floorChar {
			n.addChild(runNode(pf))
		}
	}
	if nl0 < len(l) {
		pb := l[nl0+1][nl1][nl2]
		if pb.value == m.floorChar {
			n.addChild(runNode(pb))
		}
	}

}

func (r *runner) checkSpace(n *rNode) {

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
			l2 = l[2]
		)
		if start != l && end != l && slices.Contains(r.shortestPath.toSlice(), l) {
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
	r.buildPath()
	return r
}
