package search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	want := 4
	got := Search(nums, target)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{4,5,6,7,0,1,2}
	target = 3
	want = -1
	got = Search(nums, target)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{1}
	target = 0
	want = -1
	got = Search(nums, target)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{1}
	target = 1
	want = 0
	got = Search(nums, target)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{3, 1}
	target = 1
	want = 1
	got = Search(nums, target)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
