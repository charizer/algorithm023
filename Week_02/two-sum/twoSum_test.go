package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	want := []int{0,1}
	got := TwoSum(nums, target)
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got[idx])
		}
	}
	nums = []int{3,2,4}
	target = 6
	want = []int{1,2}
	got = TwoSum(nums, target)
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got[idx])
		}
	}
	nums = []int{3,3}
	target = 6
	want = []int{0,1}
	got = TwoSum(nums, target)
	for idx := range got {
		if got[idx] != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got[idx])
		}
	}
}
