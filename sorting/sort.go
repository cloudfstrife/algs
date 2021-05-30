package sorting

import "bytes"

// -------------------------------------------------------------

type Sortable interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i must sort before the element with index j.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func Sorted(s Sortable) bool {
	for i := 1; i < s.Len(); i++ {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}

// -------------------------------------------------------------

type RuneSlice []rune

func (x RuneSlice) Len() int {
	return len(x)
}

func (x RuneSlice) Less(i, j int) bool {
	return x[i] < x[j]
}

func (x RuneSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x RuneSlice) String() string {
	buf := new(bytes.Buffer)
	for i := 0; i < x.Len(); i++ {
		buf.WriteString(string(x[i]))
		buf.WriteString(" ")
	}
	return buf.String()
}
