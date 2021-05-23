package union

// QuickUnionUF quick union union find
// Page number - Chinese Edition : 141
// Page number : 224
type QuickUnionUF struct {
	id    []int
	count int
}

// NewQuickUnionUF create QuickUnionUF
func NewQuickUnionUF(n int) *QuickUnionUF {
	uf := QuickUnionUF{
		id:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return &uf
}

// Union union
func (uf QuickUnionUF) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.id[pRoot] = qRoot
	uf.count = uf.count - 1
}

// Find find
func (uf QuickUnionUF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

// Contected return p to q is connected
func (uf QuickUnionUF) Contected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count return count union
func (uf QuickUnionUF) Count() int {
	return uf.count
}
