package search_a_2d_matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	want := true
	got := SearchMatrix(matrix, target)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target = 13
	want = false
	got = SearchMatrix(matrix, target)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
