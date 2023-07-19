package modules

type point [2]int

type floor [][]node
type dungon []floor
type layout interface {
	floor | dungon
}
type path map[point]bool

func (p path) add(s point) {
	p[s] = true
}

func (p path) update(s []point) {
	for _, x := range s {
		p[x] = true
	}
}

func (p path) setShortestPath(comp path) {
	if len(p) > len(comp) {
		p = comp
	}
}
