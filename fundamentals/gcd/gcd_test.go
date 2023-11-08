package gcd

import "testing"

func TestGCD(t *testing.T) {
	cases := map[string]struct {
		i, j, expect int
	}{
		"zero":    {i: 0, j: 0, expect: 0},
		"generic": {i: 128, j: 64, expect: 64},
		"reverse": {i: 64, j: 128, expect: 64},
	}

	for name, item := range cases {
		t.Run(name, func(t *testing.T) {
			result := GCD(item.i, item.j)
			if result != item.expect {
				t.Errorf("%s : input : %#v , %#v Want : %#v Got : %#v", name, item.i, item.j, item.expect, result)
				return
			}
		})
	}
}
