package shell

import (
	"github.com/cloudfstrife/algs/sorting"
)

// Sort shell sort
func Sort(a sorting.Sortable) {
	var h int = 1
	for h < a.Len()/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < a.Len(); i++ {
			for j := i; j >= h; j -= h {
				if a.Less(j, j-h) {
					a.Swap(j, j-h)
				}
			}
		}
		h = h / 3
	}
}
