package minimum_path_sum

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	want := 7
	got := MinPathSum(grid)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
