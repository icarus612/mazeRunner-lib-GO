package modules

import "fmt"

type point [3]int

type floor [][]node
type layout []floor
type lFuncN func(n node)
type lFunc func()

func (l layout) traverse(f1 lFuncN, f2 ...lFunc) {
	for _, x := range l {
		for _, y := range x {
			for _, z := range y {
				f1(z)
			}
			if len(f2) > 0 {
				f2[0]()
			}
		}
		if len(f2) > 1 {
			f2[1]()
		}
	}
}

func (l layout) print() {
	nl := func() { fmt.Println() }
	l.traverse(
		func(y node) {
			fmt.Print(string(y.value))
		},
		nl,
		nl,
	)
}

func (l layout) deepCopy() layout {
	nl := make(layout, len(l))
	for i, f := range l {
		nl[i] = make(floor, len(f))
		for j, c := range f {
			nl[i][j] = make([]node, len(c))
			copy(nl[i][j], l[i][j])
		}
	}
	return nl
}

//type floor [][]node
//type dungon []floor
//type layout interface {
//	floor | dungon
//}
