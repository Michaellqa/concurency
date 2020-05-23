package gotour

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestSame(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(2)
	if !Same(t1, t2) {
		t.Error("trees have different elements")
	}
}
