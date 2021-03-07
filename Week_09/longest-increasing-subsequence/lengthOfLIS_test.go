package longest_increasing_subsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	want := 4
	got := LengthOfLIS(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{0, 1, 0, 3, 2, 3}
	want = 4
	got = LengthOfLIS(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{7, 7, 7, 7, 7, 7, 7}
	want = 1
	got = LengthOfLIS(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
