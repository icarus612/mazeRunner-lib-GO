package modules

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

func (p path) toSlice() []point {
	var s []point
	for k, v := range p {
		if v {
			s = append(s, k)
		}
	}
	return s
}
