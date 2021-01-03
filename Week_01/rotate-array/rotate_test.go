package rotate_array

import (
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	want := []int{5,6,7,1,2,3,4}
	Rotate(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
	nums = []int{1,2,3,4,5,6,7}
	k = 3
	want = []int{5,6,7,1,2,3,4}
	Rotate2(nums, k)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
}
