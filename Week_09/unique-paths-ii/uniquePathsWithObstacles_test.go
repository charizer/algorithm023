package unique_paths_ii

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	want := 2
	got := UniquePathsWithObstacles(obstacleGrid)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
