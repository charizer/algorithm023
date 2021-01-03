package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	want := []int{0,1}
	got := TwoSum(nums, target)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], g)
		}
	}
}
