package modules

type point [2]int 
type floor [][]node
type layout []floor

type path map[point]bool

func (p path) add(s point) {
	p[s] = true
}

func (p path) update(s []point) {
	for _, x := range s {
		p[x] = true
	}
}