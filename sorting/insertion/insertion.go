package insertion

import "github.com/cloudfstrife/algs/sorting"

// Sort insertion sort
func Sort(a sorting.Sortable) {
	for i := 1; i < a.Len(); i++ {
		for j := i; j > 0; j-- {
			if a.Less(j, j-1) {
				a.Swap(j, j-1)
			}
		}
	}
}
