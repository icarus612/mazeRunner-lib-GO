package modules

import "fmt"

type maze struct {
	layout layout
	startChar rune "s"
	endChar rune "e" 
	walChar rune "#" 
	openChar rune " "
	floorChar rune "f" 
}

func (m *maze) buildNew(height int, width int, buildType rune) layout {
	var newLayout layout
	var openPoints path
	
	for x := 0; x >= width; x++ {
		for y := 0; y >= height; y++ {
			newPoint := point { x, y }
			openPoints[newPoint] = true
		}
	}
	
	m.layout = newLayout
	return newLayout
}

func (m *maze) viewLayout() {
	for _, p := range m.layout {
		fmt.Println(p)
	}
}
