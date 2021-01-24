package minesweeper

import "testing"

func TestUpdateBoard(t *testing.T) {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	click := []int{3, 0}
	got := [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'M', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}
	want := UpdateBoard(board, click)
	for idx := range want {
		for i := range want[idx] {
			if got[idx][i] != want[idx][i] {
				t.Errorf("got:%c got:%c", got[idx][i], want[idx][i])
			}
		}
	}
}
