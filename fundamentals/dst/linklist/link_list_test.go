package linklist

import (
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewNode(t *testing.T) {
	tests := map[string]struct {
		v      interface{}
		expect Node
	}{
		"general": {v: "Go", expect: Node{Val: "Go"}},
		"nil":     {v: nil, expect: Node{Val: nil}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewNode(tc.v)
			if !cmp.Equal(*got, tc.expect) {
				t.Errorf("NewNode(%#v) expect : %#v got : %#v ", tc.v, tc.expect, *got)
				return
			}
		})
	}
}

func TestNewLinkList(t *testing.T) {
	tests := map[string]struct {
		expect LinkList
	}{
		"general": {expect: LinkList{lock: &sync.Mutex{}, Head: nil, End: nil, Length: 0}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewLinkList()
			opts := []cmp.Option{
				cmpopts.IgnoreUnexported(tc.expect, *got),
			}
			if !cmp.Equal(*got, tc.expect, opts...) {
				t.Errorf("NewLinkList expect : %#v got : %#v ", tc.expect, *got)
				return
			}
		})
	}
}

func TestLinkListAddToHead(t *testing.T) {
	tests := map[string]struct {
		expect []interface{}
	}{
		"general":        {expect: []interface{}{""}},
		"multiple":       {expect: []interface{}{"1", "2", "3"}},
		"different type": {expect: []interface{}{"1", 2, "3"}},
		"with nil":       {expect: []interface{}{"1", nil, 2, "3"}},
		"first nil":      {expect: []interface{}{nil, "1", 2, "3"}},
		"last nil":       {expect: []interface{}{"1", 2, "3", nil}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			link := NewLinkList()
			for _, item := range tc.expect {
				link.AddToHead(item)
			}
			node := link.Head
			for i := 0; i < link.Length; i++ {
				if !cmp.Equal(node.Val, tc.expect[len(tc.expect)-1-i]) {
					t.Errorf("LinkList.AddToHead expect : %#v got : %#v ", tc.expect[len(tc.expect)-1-i], node.Val)
					return
				}
				node = node.Next
			}
		})
	}
}

func TestLinkListAddToEnd(t *testing.T) {
	tests := map[string]struct {
		expect []interface{}
	}{
		"general":        {expect: []interface{}{""}},
		"multiple":       {expect: []interface{}{"1", "2", "3"}},
		"different type": {expect: []interface{}{"1", 2, "3"}},
		"with nil":       {expect: []interface{}{"1", nil, 2, "3"}},
		"first nil":      {expect: []interface{}{nil, "1", 2, "3"}},
		"last nil":       {expect: []interface{}{"1", 2, "3", nil}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			link := NewLinkList()
			for _, item := range tc.expect {
				link.AddToEnd(item)
			}
			node := link.Head
			for i := 0; i < link.Length; i++ {
				if !cmp.Equal(node.Val, tc.expect[i]) {
					t.Errorf("LinkList.AddToEnd expect : %#v got : %#v ", tc.expect[i], node.Val)
					return
				}
				node = node.Next
			}
		})
	}
}

func TestLinkListRemoveFromHead(t *testing.T) {
	type exp struct {
		v   interface{}
		err error
	}
	tests := map[string]struct {
		put    []interface{}
		expect exp
	}{
		"general":        {put: []interface{}{""}, expect: exp{v: "", err: nil}},
		"multiple":       {put: []interface{}{"1", "2", "3"}, expect: exp{v: "1", err: nil}},
		"different type": {put: []interface{}{"1", 2, "3"}, expect: exp{v: "1", err: nil}},
		"with nil":       {put: []interface{}{"1", nil, 2, "3"}, expect: exp{v: "1", err: nil}},
		"first nil":      {put: []interface{}{nil, "1", 2, "3"}, expect: exp{v: nil, err: nil}},
		"empty":          {put: nil, expect: exp{v: nil, err: ErrEmpty}},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			link := NewLinkList()
			if tc.put != nil {
				for _, item := range tc.put {
					link.AddToEnd(item)
				}
			}
			v, err := link.RemoveFromHead()

			if !cmp.Equal(v, tc.expect.v) || !cmp.Equal(err, tc.expect.err, opt) {
				t.Errorf("LinkList.RemoveFromHead expect : %#v , %#v got : %#v , %#v", tc.expect.v, tc.expect.err, v, err)
				return
			}
		})
	}
}

func TestLinkListRemoveFromEnd(t *testing.T) {
	type exp struct {
		v   interface{}
		err error
	}
	tests := map[string]struct {
		put    []interface{}
		expect exp
	}{
		"general":        {put: []interface{}{""}, expect: exp{v: "", err: nil}},
		"multiple":       {put: []interface{}{"1", "2", "3"}, expect: exp{v: "3", err: nil}},
		"different type": {put: []interface{}{"1", 2}, expect: exp{v: 2, err: nil}},
		"with nil":       {put: []interface{}{"1", nil, 2, "3"}, expect: exp{v: "3", err: nil}},
		"first nil":      {put: []interface{}{"1", 2, "3", nil}, expect: exp{v: nil, err: nil}},
		"empty":          {put: nil, expect: exp{v: nil, err: ErrEmpty}},
	}
	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			link := NewLinkList()
			if tc.put != nil {
				for _, item := range tc.put {
					link.AddToEnd(item)
				}
			}
			v, err := link.RemoveFromEnd()
			if !cmp.Equal(v, tc.expect.v) || !cmp.Equal(err, tc.expect.err, opt) {
				t.Errorf("LinkList.RemoveFromEnd expect : %#v , %#v got : %#v , %#v", tc.expect.v, tc.expect.err, v, err)
				return
			}
		})
	}

}

func TestLinkListIsEmpty(t *testing.T) {
	tests := map[string]struct {
		put    []interface{}
		expect bool
	}{
		"general":        {put: []interface{}{""}, expect: false},
		"multiple":       {put: []interface{}{"1", "2", "3"}, expect: false},
		"different type": {put: []interface{}{"1", 2}, expect: false},
		"with nil":       {put: []interface{}{"1", nil, 2, "3"}, expect: false},
		"first nil":      {put: []interface{}{"1", 2, "3", nil}, expect: false},
		"empty":          {put: nil, expect: true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			link := NewLinkList()
			if tc.put != nil {
				for _, item := range tc.put {
					link.AddToEnd(item)
				}
			}
			v := link.IsEmpty()
			if !cmp.Equal(v, tc.expect) {
				t.Errorf("LinkList.IsEmpty expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}

func TestLinkListSize(t *testing.T) {
	tests := map[string]struct {
		put    []interface{}
		expect int
	}{
		"general":        {put: []interface{}{""}, expect: 1},
		"multiple":       {put: []interface{}{"1", "2", "3"}, expect: 3},
		"different type": {put: []interface{}{"1", 2}, expect: 2},
		"with nil":       {put: []interface{}{"1", nil, 2, "3"}, expect: 4},
		"first nil":      {put: []interface{}{"1", 2, "3", nil}, expect: 4},
		"empty":          {put: nil, expect: 0},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			link := NewLinkList()
			if tc.put != nil {
				for _, item := range tc.put {
					link.AddToEnd(item)
				}
			}
			v := link.Size()
			if !cmp.Equal(v, tc.expect) {
				t.Errorf("LinkList.Size expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}
