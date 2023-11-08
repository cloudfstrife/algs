package stack

import (
	"testing"

	"github.com/cloudfstrife/algs/fundamentals/dst/linklist"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewStack(t *testing.T) {
	tests := map[string]struct {
		expect Stack
	}{"general": {expect: Stack{link: linklist.NewLinkList()}}}

	for name, c := range tests {
		t.Run(name, func(t *testing.T) {
			queue := NewStack()
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(*queue, c.expect),
			}
			if !cmp.Equal(*queue, c.expect, opts...) {
				t.Errorf("NewStack() want : %#v got : %#v", c.expect, queue)
				return
			}
		})
	}
}

func TestStackPush(t *testing.T) {
	tests := map[string]struct {
		expect []interface{}
	}{
		"general":  {expect: []interface{}{"1", 2, "3", 4, "5"}},
		"with nil": {expect: []interface{}{nil, "1", nil, "3", 4, nil}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			stack := NewStack()
			for _, v := range tc.expect {
				stack.Push(v)
			}

			for i := 0; i < stack.Size(); i++ {
				v, err := stack.link.RemoveFromHead()
				if !cmp.Equal(tc.expect[i], v) || err != nil {
					t.Errorf("stack.Push() expect : %#v got : %#v,%#v ", tc.expect[i], v, err)
					return
				}
			}
		})
	}
}

func TestStackPop(t *testing.T) {
	type exp struct {
		v   interface{}
		err error
	}
	tests := map[string]struct {
		input  []interface{}
		expect exp
	}{
		"general":        {input: []interface{}{""}, expect: exp{v: "", err: nil}},
		"multiple":       {input: []interface{}{"1", "2", "3"}, expect: exp{v: "3", err: nil}},
		"different type": {input: []interface{}{"1", 2}, expect: exp{v: 2, err: nil}},
		"with nil":       {input: []interface{}{"1", nil, 2, "3"}, expect: exp{v: "3", err: nil}},
		"first nil":      {input: []interface{}{nil, "1", 2, "3"}, expect: exp{v: "3", err: nil}},
		"empty":          {input: nil, expect: exp{v: nil, err: linklist.ErrEmpty}},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			stack := NewStack()
			if tc.input != nil {
				for _, v := range tc.input {
					stack.Push(v)
				}
			}
			v, err := stack.Pop()
			if !cmp.Equal(tc.expect.v, v) || !cmp.Equal(tc.expect.err, err, opt) {
				t.Errorf("Stack.Pop() expect : %#v , %#v got : %#v , %#v", tc.expect.v, tc.expect.err, v, err)
				return
			}
		})
	}
}

func TestStackIsEmpty(t *testing.T) {
	tests := map[string]struct {
		input  []interface{}
		expect bool
	}{
		"empty": {input: nil, expect: true},
		"full":  {input: []interface{}{nil, "1", nil, "3", 4, nil}, expect: false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			stack := NewStack()
			if tc.input != nil {
				for _, v := range tc.input {
					stack.Push(v)
				}
			}
			v := stack.IsEmpty()
			if !cmp.Equal(tc.expect, v) {
				t.Errorf("Stack.IsEmpty() expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}

func TestStackSize(t *testing.T) {
	tests := map[string]struct {
		input  []interface{}
		expect int
	}{
		"empty": {input: nil, expect: 0},
		"full":  {input: []interface{}{nil, "1", nil, "3", 4, nil}, expect: 6},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			stack := NewStack()
			if tc.input != nil {
				for _, v := range tc.input {
					stack.Push(v)
				}
			}
			v := stack.Size()
			if !cmp.Equal(tc.expect, v) {
				t.Errorf("Stack.Size() expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}
