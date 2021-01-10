package top_k_frequent_elements

import "testing"

func TestTopKFrequent(t *testing.T) {
	nums := []int{1,1,1,2,2,3}
	k := 2
	want := []int{1,2}
	got := TopKFrequent(nums, k)
	for idx := range want {
		if got[idx] != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got[idx])
		}
	}
}
