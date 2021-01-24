package jump_game

import "testing"

func TestCanJump(t *testing.T) {
	nums := []int{2,3,1,1,4}
	want := true
	got := CanJump(nums)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
