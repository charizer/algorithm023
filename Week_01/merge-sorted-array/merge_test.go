package merge_sorted_array

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	Merge(nums1,3, nums2, 3)
	want := []int{1,2,2,3,5,6}
	for idx, n := range want {
		if nums1[idx] != n {
			t.Errorf("want:%d got:%d", n, nums1[idx])
		}
	}
}
