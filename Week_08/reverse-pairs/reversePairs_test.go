package reverse_pairs

import "testing"

func TestReversePairs(t *testing.T) {
	nums := []int{1,3,2,3,1}
	want := 2
	got := ReversePairs(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	nums = []int{2,4,3,5,1}
	want = 3
	got = ReversePairs(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
