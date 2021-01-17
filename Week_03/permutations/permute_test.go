package permutations

import "testing"

func TestPermute(t *testing.T) {
	nums := []int{1,2,3}
	want := [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,2,1},{3,1,2}}
	got := Permute(nums)
	for idx := range got {
		for i := range got[idx] {
			if got[idx][i] != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], got[idx][i])
			}
		}
	}
}
