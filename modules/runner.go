package modules

import (
	"fmt"
)

type runner struct {
	completed     bool
	openNodes     []node
	possiblePaths []path
	visited       path
	toVisit       path
	start         node
	end           node
	maze          maze
	path          path
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
	func check(nodeValue point) {
		for _, x := range r.openNodes {

		}
		if nodeValue
	}
}

func (r *runner) lookAround(n node) {
	for _, v := range r.openNodes {
		if v.location[0]-1 == n.location[0] && v.location[1] == n.location[1] {
			n.addChild(v)
		} else if v.location[0]+1 == n.location[0] && v.location[1] == n.location[1] {
			n.addChild(v)
		} else if v.location[1]-1 == n.location[1] && v.location[0] == n.location[0] {
			n.addChild(v)
		} else if v.location[1]+1 == n.location[1] && v.location[0] == n.location[0] {
			n.addChild(v)
		}
	}
}

func (r *runner) makeNodePaths() {
	r.toVisit.add(r.start)
}

func (r runner) viewCompleted() {
	for _, row := range r.mappedMaze {
		fmt.Println(row)
	}
}

func (r *runner) buildPath(path) {

}

func (r *runner) setShortestPath(p path) {
	if len(p) < len(r.path) {
		r.path = p
	}
}

func (r runner) Runner(m maze) runner {
	r.completed = false
	r.maze = m
	r.getOpenNodes()
	r.findEndPoints()
	return r
}
