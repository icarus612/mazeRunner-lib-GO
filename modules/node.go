package modules

type node struct {
	value    rune
	location point
}
type rNode struct {
	node
	path     []point
	children []rNode
}

func (n *rNode) addChild(c rNode) {
	n.children = append(n.children, c)
}

func (n *rNode) removeChild(c rNode) {
	n.children = append(n.children, c)
}

func (r *rNode) setShortestPath(p []point) {
	if len(p) < len(r.path) {
		r.path = p
	}
}
