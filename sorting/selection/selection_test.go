package selection

import (
	"os"
	"testing"

	"github.com/cloudfstrife/algs/commons"
	"github.com/cloudfstrife/algs/sorting"
)

const inPath = "../../data/algs4-data/tiny.txt"

func TestSelectSort(t *testing.T) {
	var err error
	var inFile *os.File
	if inFile, err = os.Open(inPath); err != nil {
		t.Errorf("open file failed : %v", err)
		return
	}
	defer inFile.Close()
	in := commons.NewIn(inFile)

	for !in.IsEmpty() {
		var all []rune = in.ReadAllRune()
		t.Log(sorting.RuneSlice(all))
		Sort(sorting.RuneSlice(all))
		t.Log(sorting.RuneSlice(all))
	}
}
