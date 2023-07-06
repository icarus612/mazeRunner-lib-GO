package modules

type mazeNode struct {
	value rune
	location point
	children []node
}

func (n *node) addChild(c node) {
	n.children = append(n.children, c)
}

func (n *node) removeChild(c node) {
	n.children = append(n.children, c)
}


type runnerNode {
	mazeNode,
	path 
}

func (n *node) setShortestPath(p path) {
	if len(p) < len(n.path) {
		n.path = p
	}
}
