package bag

import (
	"testing"

	"github.com/cloudfstrife/algs/fundamentals/dst/linklist"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewBag(t *testing.T) {
	tests := map[string]struct {
		expect Bag
	}{
		"general": {Bag{link: linklist.NewLinkList()}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			bag := NewBag()
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(*bag, tc.expect),
			}
			if !cmp.Equal(*bag, tc.expect, opts...) {
				t.Errorf("NewBag expect : %#v got : %#v ", tc.expect, *bag)
				return
			}
		})
	}
}

func TestBagAdd(t *testing.T) {
	tests := map[string]struct {
		expect []interface{}
	}{
		"general":  {expect: []interface{}{"1", 2, "3", 4, "5"}},
		"with nil": {expect: []interface{}{nil, "1", nil, "3", 4, nil}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			bag := NewBag()
			for _, v := range tc.expect {
				bag.Add(v)
			}

			for i := 0; i < bag.Size(); i++ {
				v, err := bag.link.RemoveFromHead()
				if !cmp.Equal(tc.expect[i], v) || err != nil {
					t.Errorf("Bag.Add expect : %#v got : %#v,%#v ", tc.expect[i], v, err)
					return
				}
			}
		})
	}
}

func TestBagIsEmpty(t *testing.T) {
	tests := map[string]struct {
		input  []interface{}
		expect bool
	}{
		"empty": {input: nil, expect: true},
		"full":  {input: []interface{}{nil, "1", nil, "3", 4, nil}, expect: false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			bag := NewBag()
			if tc.input != nil {
				for _, v := range tc.input {
					bag.Add(v)
				}
			}

			v := bag.IsEmpty()
			if !cmp.Equal(tc.expect, v) {
				t.Errorf("Bag.IsEmpty expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}

func TestBagSize(t *testing.T) {
	tests := map[string]struct {
		input  []interface{}
		expect int
	}{
		"empty": {input: nil, expect: 0},
		"full":  {input: []interface{}{nil, "1", nil, "3", 4, nil}, expect: 6},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			bag := NewBag()
			if tc.input != nil {
				for _, v := range tc.input {
					bag.Add(v)
				}
			}
			v := bag.Size()
			if !cmp.Equal(tc.expect, v) {
				t.Errorf("Bag.Size expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}
