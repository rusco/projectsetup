package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"my.com/calc"
)

func TestMain(t *testing.T) {
	t.Parallel()
	var (
		want1 = 5
		got1  = calc.Add(2, 3)
	)

	if !cmp.Equal(want1, got1) {
		t.Error(cmp.Diff(want1, got1))
	}

}
