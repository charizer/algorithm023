package find_minimum_in_rotated_sorted_array

import "testing"

func TestFindMin(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	want := 1
	got := FindMin(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	want = 0
	got = FindMin(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{1}
	want = 1
	got = FindMin(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}

	nums = []int{11, 13, 15, 17}
	want = 11
	got = FindMin(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
