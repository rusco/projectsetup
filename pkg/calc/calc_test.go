package calc_test

import (
	"testing"

	calc "my.com/calc"

	"github.com/google/go-cmp/cmp"
)

func TestSimple(t *testing.T) {
	t.Parallel()
	var (
		want1 = 5
		got1  = calc.Add(2, 3)
		want2 = 5
		got2  = calc.Sub(10, 5)
		want3 = 10
		got3  = calc.Mult(5, 2)
		want4 = 5
		got4  = calc.Div(25, 5)
	)

	if !cmp.Equal(want1, got1) {
		t.Error(cmp.Diff(want1, got1))
	}
	if !cmp.Equal(want2, got2) {
		t.Error(cmp.Diff(want2, got2))
	}
	if !cmp.Equal(want3, got3) {
		t.Error(cmp.Diff(want3, got3))
	}
	if !cmp.Equal(want4, got4) {
		t.Error(cmp.Diff(want4, got4))
	}

}
