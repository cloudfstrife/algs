package find

// QuickFindUF quick find union find
// Page number - Chinese Edition : 139-140
// Page number : 221-222
type QuickFindUF struct {
	id    []int
	count int
}

// NewQuickFindUF create QuickFindUF
func NewQuickFindUF(n int) *QuickFindUF {
	uf := QuickFindUF{
		id:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return &uf
}

// Union union
func (uf QuickFindUF) Union(p, q int) {
	pv := uf.Find(p)
	qv := uf.Find(q)
	if pv == qv {
		return
	}
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pv {
			uf.id[i] = qv
		}
	}
	uf.count = uf.count - 1
}

// Find find
func (uf QuickFindUF) Find(p int) int {
	return uf.id[p]
}

// Contected return p to q is connected
func (uf QuickFindUF) Contected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count return count union
func (uf QuickFindUF) Count() int {
	return uf.count
}
