package jump_game_ii

import "testing"

func TestJump(t *testing.T) {
	nums := []int{2,3,1,1,4}
	want := 2
	got := Jump(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
