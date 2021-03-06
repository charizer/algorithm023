package frog_jump

import "testing"

func TestCanCross(t *testing.T) {
	stones := []int{0,1,3,5,6,8,12,17}
	want := true
	got := CanCross(stones)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	stones = []int{0,1,2,3,4,8,9,11}
	want = false
	got = CanCross(stones)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
