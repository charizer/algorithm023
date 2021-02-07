package max_sum_of_rectangle_no_larger_than_k

import "testing"

func TestMaxSumSubmatrix(t *testing.T) {
	matrix := [][]int{{1, 0, 1}, {0, -2, 3}}
	k := 2
	want := 2
	got := MaxSumSubmatrix(matrix, k)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
