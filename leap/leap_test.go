package leap

import "testing"

type lp struct {
	year int
	isleap bool
}

var list = []lp {
	lp{1900, false},
	lp{1996, true},
	lp{1997, false},
	lp{2000, true},
	lp{2001, false},
}

func Test_Basic(t *testing.T) {
	for _, j := range(list) {
		if Isleap(j.year) != j.isleap {
			t.Errorf("FAIL: %v\n", j)
		}
	}
}
