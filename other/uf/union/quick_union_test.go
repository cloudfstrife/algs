package union

import (
	"os"
	"testing"

	"github.com/cloudfstrife/algs/commons"
)

const inPath = "../../../data/algs4-data/mediumUF.txt"

func TestUF(t *testing.T) {
	var err error
	var inFile *os.File
	if inFile, err = os.Open(inPath); err != nil {
		t.Errorf("open file failed : %v", err)
		return
	}
	defer inFile.Close()
	in := commons.NewIn(inFile)
	var n int
	n, err = in.ReadInteger()
	if err != nil {
		t.Errorf("read input file failed : %v", err)
		return
	}
	u := NewQuickUnionUF(n)

	for !in.IsEmpty() {
		var p, q int
		p, err = in.ReadInteger()
		if err != nil {
			t.Errorf("read input file failed : %v", err)
			return
		}
		q, err = in.ReadInteger()
		if err != nil {
			t.Errorf("read input file failed : %v", err)
			return
		}
		if u.Contected(p, q) {
			continue
		}
		u.Union(p, q)
		t.Logf("%d %d", p, q)
	}
}
