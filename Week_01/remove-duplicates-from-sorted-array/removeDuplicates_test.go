package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1,1,2}
	want := 2
	got := RemoveDuplicates(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{0,0,1,1,1,2,2,3,3,4}
	want = 5
	got = RemoveDuplicates(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
