package commons

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewIn(t *testing.T) {
	testCases := map[string]struct {
		reader io.Reader
		want   string
	}{
		"string": {
			reader: strings.NewReader("You\nAre\nNot\nAlone"),
			want:   "You",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.ReadString()
			if v != tc.want {
				t.Errorf("%s : Want : %#v Got : %#v", name, tc.want, v)
				return
			}
		})
	}
}

func TestInIsEmpty(t *testing.T) {
	testCases := map[string]struct {
		reader io.Reader
		want   bool
	}{
		"full": {
			reader: strings.NewReader("You\nAre\nNot\nAlone"),
			want:   false,
		},
		"empty": {
			reader: strings.NewReader(""),
			want:   true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.IsEmpty()
			if v != tc.want {
				t.Errorf("%s : Want : %#v Got : %#v", name, tc.want, v)
				return
			}
		})
	}
}

func TestInReadString(t *testing.T) {
	testCases := map[string]struct {
		reader io.Reader
		want   string
	}{
		"full": {
			reader: strings.NewReader("You \nAre\nNot\nAlone"),
			want:   "You",
		},
		"empty": {
			reader: strings.NewReader(""),
			want:   "",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.ReadString()
			if v != tc.want {
				t.Errorf("%s : Want : %#v Got : %#v", name, tc.want, v)
				return
			}
		})
	}
}

func TestInReadAllString(t *testing.T) {
	testCases := map[string]struct {
		reader io.Reader
		want   []string
	}{
		"full": {
			reader: strings.NewReader("You\n Are\nNot  \n Alone  "),
			want:   []string{"You", "Are", "Not", "Alone"},
		},
		"empty": {
			reader: strings.NewReader(""),
			want:   nil,
		},
		"space": {
			reader: strings.NewReader(" \nYou"),
			want:   []string{"You"},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.ReadAllString()
			if !cmp.Equal(v, tc.want) {
				t.Errorf("%s : Want : %#v Got : %#v", name, tc.want, v)
				return
			}
		})
	}
}

func TestInReadInteger(t *testing.T) {
	type wantT struct {
		i   int
		err error
	}
	testCases := map[string]struct {
		reader io.Reader
		want   wantT
	}{
		"general": {
			reader: strings.NewReader("9 8\n7\n6\n5\n4\n3\n2\n1"),
			want: wantT{
				i:   9,
				err: nil,
			},
		},
		"with-space": {
			reader: strings.NewReader("9 \n 8 \n 7\n6  \n  5\n  4  \n    3    \n2    \n1"),
			want: wantT{
				i:   9,
				err: nil,
			},
		},
		"empty": {
			reader: strings.NewReader(""),
			want: wantT{
				i:   0,
				err: &strconv.NumError{Func: "Atoi", Num: "", Err: strconv.ErrSyntax},
			},
		},
		"Not-Int": {
			reader: strings.NewReader("alone\ns"),
			want: wantT{
				i:   0,
				err: &strconv.NumError{Func: "Atoi", Num: "alone", Err: strconv.ErrSyntax},
			},
		},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v, err := in.ReadInteger()
			if !cmp.Equal(v, tc.want.i) || !cmp.Equal(tc.want.err, err, opt) {
				t.Errorf("%s : Want : %#v , %#v  Got : %#v , %#v", name, tc.want.i, tc.want.err, v, err)
				return
			}
		})
	}
}

func TestInReadAllInteger(t *testing.T) {
	testCases := map[string]struct {
		reader io.Reader
		want   []int
	}{
		"general": {
			reader: strings.NewReader("9\n8\n7\n6\n5\n4\n3\n2\n1"),
			want:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		"with-space": {
			reader: strings.NewReader("9 \n 8 \n 7\n6  \n  5\n  4  \n    3    \n2    \n1"),
			want:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		"empty": {
			reader: strings.NewReader(""),
			want:   []int{},
		},
		"Not-Int": {
			reader: strings.NewReader("alone\ns"),
			want:   []int{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.ReadAllInteger()
			if !cmp.Equal(v, tc.want) {
				t.Errorf("%s : Want :  %#v  Got : %#v", name, tc.want, v)
				return
			}
		})
	}
}

func TestInReadBool(t *testing.T) {
	type wantT struct {
		b   bool
		err error
	}
	testCases := map[string]struct {
		reader io.Reader
		want   wantT
	}{
		"general-true": {
			reader: strings.NewReader("true\n"),
			want: wantT{
				b:   true,
				err: nil,
			},
		},
		"general-1": {
			reader: strings.NewReader("1\n"),
			want: wantT{
				b:   true,
				err: nil,
			},
		},
		"general-false": {
			reader: strings.NewReader("false\n"),
			want: wantT{
				b:   false,
				err: nil,
			},
		},
		"general-0": {
			reader: strings.NewReader("0\n"),
			want: wantT{
				b:   false,
				err: nil,
			},
		},
		"with-space": {
			reader: strings.NewReader("  0  \n"),
			want: wantT{
				b:   false,
				err: nil,
			},
		},
		"empty": {
			reader: strings.NewReader(""),
			want: wantT{
				b:   false,
				err: fmt.Errorf("attempts to read a 'bool' value , but the next token is %#v ", ""),
			},
		},
		"Not-Bool": {
			reader: strings.NewReader("alone\ns"),
			want: wantT{
				b:   false,
				err: fmt.Errorf("attempts to read a 'bool' value , but the next token is %#v ", "alone"),
			},
		},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v, err := in.ReadBool()
			if !cmp.Equal(v, tc.want.b) || !cmp.Equal(tc.want.err, err, opt) {
				t.Errorf("%s : Want : %#v , %#v  Got : %#v , %#v", name, tc.want.b, tc.want.err, v, err)
				return
			}
		})
	}
}

func TestInReadAllBool(t *testing.T) {
	testCases := map[string]struct {
		reader io.Reader
		want   []bool
	}{
		"general": {
			reader: strings.NewReader("true\n1\nfalse\n0"),
			want:   []bool{true, true, false, false},
		},
		"with-space": {
			reader: strings.NewReader("true \n 1 \nfalse \n   0"),
			want:   []bool{true, true, false, false},
		},
		"empty": {
			reader: strings.NewReader(""),
			want:   []bool{},
		},
		"Not-Bool": {
			reader: strings.NewReader("alone\n"),
			want:   []bool{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.ReadAllBool()
			if !cmp.Equal(v, tc.want) {
				t.Errorf("%s : Want : %#v   Got : %#v ", name, tc.want, v)
				return
			}
		})
	}
}

func TestInReadRune(t *testing.T) {
	type wantT struct {
		r   rune
		i   int
		err error
	}
	testCases := map[string]struct {
		reader io.Reader
		want   wantT
	}{
		"general-multi": {
			reader: strings.NewReader("true\n"),
			want: wantT{
				r:   't',
				i:   1,
				err: nil,
			},
		},
		"general-one": {
			reader: strings.NewReader("1\n"),
			want: wantT{
				r:   '1',
				i:   1,
				err: nil,
			},
		},
		"general-chinese": {
			reader: strings.NewReader("你好\n"),
			want: wantT{
				r:   '你',
				i:   3,
				err: nil,
			},
		},
		"empty": {
			reader: strings.NewReader(""),
			want: wantT{
				r:   0,
				i:   0,
				err: errors.New("internal In.ReadRune() error"),
			},
		},
		"Not-rune": {
			reader: strings.NewReader(string([]byte{0xD8, 0x00})),
			want: wantT{
				r:   0,
				i:   0,
				err: errors.New("internal In.ReadRune() error"),
			},
		},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v, len, err := in.ReadRune()
			if !cmp.Equal(v, tc.want.r) || !cmp.Equal(len, tc.want.i) || !cmp.Equal(tc.want.err, err, opt) {
				t.Errorf("%s : Want : %#v , %#v , %#v  Got : %#v , %#v , %#v", name, tc.want.r, tc.want.i, tc.want.err, v, len, err)
				return
			}
		})
	}
}

func TestInReadAllRune(t *testing.T) {
	testCases := map[string]struct {
		reader io.Reader
		want   []rune
	}{
		"general-multi": {
			reader: strings.NewReader("true\n,你好"),
			want:   []rune{'t', 'r', 'u', 'e', ',', '你', '好'},
		},
		"empty": {
			reader: strings.NewReader(""),
			want:   []rune{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.ReadAllRune()
			if !cmp.Equal(v, tc.want) {
				t.Errorf("%s : Want : %#v Got : %#v ", name, tc.want, v)
				return
			}
		})
	}
}

func TestInReadFloat(t *testing.T) {
	type wantT struct {
		f   float64
		err error
	}
	testCases := map[string]struct {
		reader io.Reader
		want   wantT
	}{
		"general": {
			reader: strings.NewReader("9.12\n85.32\n7.12"),
			want: wantT{
				f:   9.12,
				err: nil,
			},
		},
		"with-space": {
			reader: strings.NewReader("9.12 \n 85.32 \n7.12 "),
			want: wantT{
				f:   9.12,
				err: nil,
			},
		},
		"empty": {
			reader: strings.NewReader(""),
			want: wantT{
				f:   0.00,
				err: &strconv.NumError{Func: "ParseFloat", Num: "", Err: strconv.ErrSyntax},
			},
		},
		"Not-float": {
			reader: strings.NewReader("alone\ns"),
			want: wantT{
				f:   0.00,
				err: &strconv.NumError{Func: "ParseFloat", Num: "alone", Err: strconv.ErrSyntax},
			},
		},
	}

	opt := cmp.Comparer(func(x, y error) bool {
		return (x == nil && y == nil) || ((x != nil && y != nil) && x.Error() == y.Error())
	})

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v, err := in.ReadFloat()
			if !cmp.Equal(v, tc.want.f) || !cmp.Equal(tc.want.err, err, opt) {
				t.Errorf("%s : Want : %#v , %#v  Got : %#v , %#v", name, tc.want.f, tc.want.err, v, err)
				return
			}
		})
	}
}

func TestInReadAllFloat(t *testing.T) {

	testCases := map[string]struct {
		reader io.Reader
		want   []float64
	}{
		"general": {
			reader: strings.NewReader("9.12\n85.32\n7.12"),
			want:   []float64{9.12, 85.32, 7.12},
		},
		"with-space": {
			reader: strings.NewReader("9.12 \n 85.32 \n7.12 "),
			want:   []float64{9.12, 85.32, 7.12},
		},
		"empty": {
			reader: strings.NewReader(""),
			want:   []float64{},
		},
		"Not-float": {
			reader: strings.NewReader("alone\ns"),
			want:   []float64{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			in := NewIn(tc.reader)
			v := in.ReadAllFloat()
			if !cmp.Equal(v, tc.want) {
				t.Errorf("%s : Want : %#v  Got : %#v ", name, tc.want, v)
				return
			}
		})
	}
}
