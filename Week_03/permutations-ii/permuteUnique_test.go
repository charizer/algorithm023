package permutations_ii

import "testing"

func TestPermuteUnique(t *testing.T) {
	nums := []int{1,1,2}
	got := PermuteUnique(nums)
	want := [][]int{{1,1,2},{1,2,1}, {2,1,1}}
	for idx := range got {
		for i := range got[idx] {
			if got[idx][i] != want[idx][i] {
				t.Errorf("want:%d got:%d", want[idx][i], got[idx][i])
			}
		}
	}
}
