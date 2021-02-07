package maximal_square

import "testing"

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	want := 4
	got := MaximalSquare(matrix)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
