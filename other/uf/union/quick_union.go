package union

type QuickUnionUF struct {
	id    []int
	count int
}

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

func (uf QuickUnionUF) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.id[pRoot] = qRoot
	uf.count = uf.count - 1
}

func (uf QuickUnionUF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf QuickUnionUF) Contected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf QuickUnionUF) Count() int {
	return uf.count
}
