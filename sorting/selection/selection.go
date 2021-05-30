package selection

import "github.com/cloudfstrife/algs/sorting"

// Sort selection sort
func Sort(a sorting.Sortable) {
	for i := 0; i < a.Len(); i++ {
		min := i
		for j := i + 1; j < a.Len(); j++ {
			if a.Less(j, min) {
				min = j
			}
		}
		a.Swap(i, min)
	}
}
