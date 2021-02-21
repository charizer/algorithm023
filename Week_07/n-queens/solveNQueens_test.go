package n_queens

import "testing"

func TestSolveNQueens(t *testing.T) {
	n := 4
	got := SolveNQueens(n)
	t.Logf("got:%+v", got)
}
