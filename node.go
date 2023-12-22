package mazerunner

type node struct {
	value    rune
	location point
}
type rNode struct {
	node
	path     path
	children []rNode
}

func (n *rNode) addChild(c rNode) {
	n.children = append(n.children, c)
}

func (r *rNode) setPath(p path) {
	r.path = p
}

func runNode(n node) rNode {
	rn := rNode{
		node: n,
		path: make(path),
	}
	return rn
}
