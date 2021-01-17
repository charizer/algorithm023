package combinations

import "testing"

func TestCombine(t *testing.T) {
	n, k := 4, 2
	want := [][]int{{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}}
	got := Combine(n,k)
	for idx := range got {
		for i := range got[idx] {
			if got[idx][i] != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], got[idx][i])
			}
		}
	}
}
