package insertion

import "github.com/cloudfstrife/algs/sorting"

// Sort selection sort
func Sort(a sorting.Sortable) {
	for i := 1; i < a.Len(); i++ {
		for j := 0; j < i; j++ {
			if a.Less(i, j) {
				a.Swap(i, j)
			}
		}
	}
}
