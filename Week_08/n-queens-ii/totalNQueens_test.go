package n_queens_ii

import "testing"

func TestTotalNQueens(t *testing.T) {
	n := 4
	want := 2
	got := TotalNQueens(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
