package modules

type node struct {
	value    rune
	location point
	children []node
}

func (n *node) addChild(c node) {
	n.children = append(n.children, c)
}

func (n *node) removeChild(c node) {
	n.children = append(n.children, c)
}
