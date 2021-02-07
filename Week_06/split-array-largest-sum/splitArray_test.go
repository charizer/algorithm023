package split_array_largest_sum

import "testing"

func TestSplitArray(t *testing.T) {
	nums := []int{7,2,5,10,8}
	m := 2
	want := 18
	got := SplitArray(nums, m)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
