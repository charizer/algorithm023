package min_cost_climbing_stairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	want := 6
	got := MinCostClimbingStairs(cost)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
