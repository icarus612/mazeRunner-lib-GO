package modules

import "fmt"

type runner struct {
	completed bool
	openNodes []node
	possiblePaths []path
	visited path
	toVisit path
	start node
	end node
	maze maze
	mappedMaze layout
}

func (r *runner) getOpenNodes() {
	p := r.maze.layout
	for x := 0; x < len(p); x++ {
		for y := 0; x < len(p[x]); y++ {
			if p[x][y] != r.maze.wallChar {
				newNode := node {

				}
				r.openNodes = append(r.openNodes, newNode) 
			}		
		} 
	} 
}

func (r *runner) findEndPoints() {

}

func (r *runner) lookAround(n node) {
	for _, v := range r.openNodes {
		if v.location[0]-1 == n.location[0] && v.location[1] == n.location[1] {
			n.addChild(v)
		} else if v. location[0]+1 == n.location[0] && v.location[1] == n.location[1] {
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

func (r runner) Runner(maze layout) runner {
	r.completed = false
	r.maze = maze
	r.getOpenNodes()
	r.findEndPoints()
	return r 
}