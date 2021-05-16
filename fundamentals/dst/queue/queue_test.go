package queue

import (
	"testing"

	"github.com/cloudfstrife/algs/fundamentals/dst/linklist"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewQueue(t *testing.T) {
	tests := map[string]struct {
		expect Queue
	}{
		"general": {expect: Queue{link: linklist.NewLinkList()}},
	}

	for name, c := range tests {
		t.Run(name, func(t *testing.T) {
			queue := NewQueue()
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(*queue, c.expect),
			}
			if !cmp.Equal(*queue, c.expect, opts...) {
				t.Errorf("NewQueue() want : %#v got : %#v", c.expect, queue)
				return
			}
		})
	}
}

func TestQueueEnQueue(t *testing.T) {
	tests := map[string]struct {
		expect []interface{}
	}{
		"general":  {expect: []interface{}{"1", 2, "3", 4, "5"}},
		"with nil": {expect: []interface{}{nil, "1", nil, "3", 4, nil}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			queue := NewQueue()
			for _, v := range tc.expect {
				queue.EnQueue(v)
			}

			for i := 0; i < queue.Size(); i++ {
				v, err := queue.link.RemoveFromHead()
				if !cmp.Equal(tc.expect[i], v) || err != nil {
					t.Errorf("Queue.EnQueue() expect : %#v got : %#v,%#v ", tc.expect[i], v, err)
					return
				}
			}
		})
	}
}

func TestQueueDeQueue(t *testing.T) {
	type exp struct {
		v   interface{}
		err error
	}
	tests := map[string]struct {
		input  []interface{}
		expect exp
	}{
		"general":        {input: []interface{}{""}, expect: exp{v: "", err: nil}},
		"multiple":       {input: []interface{}{"1", "2", "3"}, expect: exp{v: "1", err: nil}},
		"different type": {input: []interface{}{"1", 2}, expect: exp{v: "1", err: nil}},
		"with nil":       {input: []interface{}{"1", nil, 2, "3"}, expect: exp{v: "1", err: nil}},
		"first nil":      {input: []interface{}{nil, "1", 2, "3"}, expect: exp{v: nil, err: nil}},
		"empty":          {input: nil, expect: exp{v: nil, err: linklist.ErrEmpty}},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			queue := NewQueue()
			if tc.input != nil {
				for _, v := range tc.input {
					queue.EnQueue(v)
				}
			}
			v, err := queue.DeQueue()
			if !cmp.Equal(tc.expect.v, v) || !cmp.Equal(tc.expect.err, err, opt) {
				t.Errorf("Queue.DeQueue() expect : %#v , %#v got : %#v , %#v", tc.expect.v, tc.expect.err, v, err)
				return
			}
		})
	}
}

func TestQueueIsEmpty(t *testing.T) {
	tests := map[string]struct {
		input  []interface{}
		expect bool
	}{
		"empty": {input: nil, expect: true},
		"full":  {input: []interface{}{nil, "1", nil, "3", 4, nil}, expect: false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			queue := NewQueue()
			if tc.input != nil {
				for _, v := range tc.input {
					queue.EnQueue(v)
				}
			}
			v := queue.IsEmpty()
			if !cmp.Equal(tc.expect, v) {
				t.Errorf("Queue.IsEmpty() expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}

func TestQueueSize(t *testing.T) {
	tests := map[string]struct {
		input  []interface{}
		expect int
	}{
		"empty": {input: nil, expect: 0},
		"full":  {input: []interface{}{nil, "1", nil, "3", 4, nil}, expect: 6},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			queue := NewQueue()
			if tc.input != nil {
				for _, v := range tc.input {
					queue.EnQueue(v)
				}
			}
			v := queue.Size()
			if !cmp.Equal(tc.expect, v) {
				t.Errorf("Queue.Size() expect : %#v got : %#v ", tc.expect, v)
				return
			}
		})
	}
}
