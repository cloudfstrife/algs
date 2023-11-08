package weighted

// WeightedQuickUnionUF weighted quick union union find
// Page number - Chinese Edition : 141
// Page number : 224
type WeightedQuickUnionUF struct {
	id    []int
	sz    []int
	count int
}

// NewWeightedQuickUnionUF create WeightedQuickUnionUF
func NewWeightedQuickUnionUF(n int) *WeightedQuickUnionUF {
	uf := WeightedQuickUnionUF{
		id:    make([]int, n),
		sz:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return &uf
}

// Union union
func (uf WeightedQuickUnionUF) Union(p, q int) {
	i := uf.Find(p)
	j := uf.Find(q)
	if i == j {
		return
	}
	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] += uf.sz[i]
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
	}
}

// Find find
func (uf WeightedQuickUnionUF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

// Contected return p to q is connected
func (uf WeightedQuickUnionUF) Contected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count return count union
func (uf WeightedQuickUnionUF) Count() int {
	return uf.count
}
