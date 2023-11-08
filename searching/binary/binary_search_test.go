package binary

import (
	"os"
	"sort"
	"testing"

	"github.com/cloudfstrife/algs/commons"
)

func TestSearch(t *testing.T) {
	cases := map[string]struct {
		i      int
		a      []int
		expect int
	}{
		"zero":       {i: 0, a: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expect: 0},
		"general-01": {i: 128, a: []int{9, 23, 52, 67, 88, 128, 564}, expect: 5},
		"general-02": {i: 128, a: []int{9, 23, 52, 67, 88, 128}, expect: 5},
		"not-exists": {i: 64, a: []int{}, expect: -1},
	}

	for name, item := range cases {
		t.Run(name, func(t *testing.T) {
			result := Search(item.i, item.a)
			if result != item.expect {
				t.Errorf("%s : input : %#v , %#v Want : %#v Got : %#v", name, item.i, item.a, item.expect, result)
				return
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	// 数据准备
	dataFilePath := "../../data/algs4-data/largeW.txt"
	var (
		file *os.File
		err  error
	)
	if file, err = os.Open(dataFilePath); err != nil {
		b.Errorf("open file %#v %#v", dataFilePath, err)
	}
	defer file.Close()
	in := commons.NewIn(file)
	data := in.ReadAllInteger()
	sort.Ints(data)

	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Search(191271, data)
	}
}

func TestRecursionSearch(t *testing.T) {
	cases := map[string]struct {
		i      int
		a      []int
		expect int
	}{
		"zero":       {i: 0, a: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expect: 0},
		"general-01": {i: 128, a: []int{9, 23, 52, 67, 88, 128, 564}, expect: 5},
		"general-02": {i: 128, a: []int{9, 23, 52, 67, 88, 128}, expect: 5},
		"not-exists": {i: 64, a: []int{}, expect: -1},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if got := RecursionSearch(c.i, c.a, 0, len(c.a)-1); got != c.expect {
				t.Errorf("RecursionSearch(%#v,%#v,%#v,%#v) = %v, want %v", c.i, c.a, 0, len(c.a)-1, got, c.expect)
				return
			}
		})
	}
}

func BenchmarkRecursionSearch(b *testing.B) {
	// 数据准备
	dataFilePath := "../../data/algs4-data/largeW.txt"
	var (
		file *os.File
		err  error
	)
	if file, err = os.Open(dataFilePath); err != nil {
		b.Errorf("open file %#v %#v", dataFilePath, err)
	}
	defer file.Close()
	in := commons.NewIn(file)
	data := in.ReadAllInteger()
	sort.Ints(data)

	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RecursionSearch(191271, data, 0, len(data)-1)
	}
}
