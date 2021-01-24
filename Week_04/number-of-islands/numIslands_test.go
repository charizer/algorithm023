package number_of_islands

import "testing"

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	want := 1
	got := NumIslands(grid)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
