package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	n := 2
	want := 2
	got := ClimbStairs(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	n = 3
	want = 3
	got = ClimbStairs(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}