package plus_one

import "testing"

func TestPlusOne(t *testing.T) {
	nums := []int{1,2,3}
	want := []int{1,2,4}
	got := PlusOne(nums)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], g)
		}
	}
}
