package modules

type point interface {
	[3]int | [2]int
}

type floor [...][...]node
type dungon [...]floor
type layout interface {
	*floor | *dungon
}
type path map[point]bool

func (p path) add(s point) {
	p[s] = true
}

func (p path) update[P point](s []P) {
	for _, x := range s {
		p[x] = true
	}
}

func (p path) setShortestPath(comp path) {
	if len(p) < len(comp) {
		p.path = comp
	}
}
