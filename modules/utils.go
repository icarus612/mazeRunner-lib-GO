package modules

import "fmt"

type point [2]int

type layout [][]node
type lFuncN func(n node)
type lFunc func()

func (l layout) traverse(f1 lFuncN, f2 ...lFunc) {
	for _, x := range l {
		for _, y := range x {
			f1(y)
		}
		if len(f2) > 0 {
			f2[0]()
		}
	}
}

func (l layout) print() {
	l.traverse(
		func(y node) {
			fmt.Print(string(y.value))
		},
		func() {
			fmt.Println()
		},
	)
}

func (l layout) deepCopy() layout {
	nl := make(layout, len(l))
	for i := range l {
		nl[i] = make([]node, len(l[i]))
		copy(nl[i], l[i])
	}
	return nl
}

//type floor [][]node
//type dungon []floor
//type layout interface {
//	floor | dungon
//}
