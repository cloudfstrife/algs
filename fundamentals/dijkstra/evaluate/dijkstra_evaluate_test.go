package evaluate

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cloudfstrife/algs/fundamentals/dst/linklist"

	"github.com/google/go-cmp/cmp"
)

func TestDijkstraEvluate(t *testing.T) {
	type exp struct {
		v   float64
		err error
	}
	tests := map[string]struct {
		evluate string
		expect  exp
	}{
		"general": {
			evluate: "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )",
			expect:  exp{v: 101, err: nil},
		},
		"with SQRT": {
			evluate: "( ( ( 1 + sqrt ( 4.0 ) ) / 2.0 ) - 1.0 )",
			expect:  exp{v: 0.5, err: nil},
		},
		"With ERROR": {
			evluate: "( 1f + 2.0f )",
			expect:  exp{v: 0, err: &strconv.NumError{Func: "ParseFloat", Num: "1f", Err: strconv.ErrSyntax}},
		},
		"With + ERROR": {
			evluate: "( + 2.0 )",
			expect:  exp{v: 0, err: linklist.ErrEmpty},
		},
		"With - ERROR": {
			evluate: "( - 2.0 )",
			expect:  exp{v: 0, err: linklist.ErrEmpty},
		},
		"With * ERROR": {
			evluate: "( * 2.0 )",
			expect:  exp{v: 0, err: linklist.ErrEmpty},
		},
		"With / ERROR": {
			evluate: "( / 2.0 )",
			expect:  exp{v: 0, err: linklist.ErrEmpty},
		},
		"Without number": {
			evluate: "( / + )",
			expect:  exp{v: 0, err: linklist.ErrEmpty},
		},
		"empty": {
			evluate: "+",
			expect:  exp{v: 0, err: linklist.ErrEmpty},
		},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			slist := strings.Split(tc.evluate, " ")
			v, err := DijkstraEvluate(slist)
			if !cmp.Equal(tc.expect.v, v) || !cmp.Equal(tc.expect.err, err, opt) {
				t.Errorf("DijkstraEvluate(%#v) expect : %#v , %#v got : %#v , %#v", slist, tc.expect.v, tc.expect.err, v, err)
				return
			}
		})
	}
}
