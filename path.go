package mazerunner

type path map[point]bool

func (p path) add(s point) {
	p[s] = true
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
