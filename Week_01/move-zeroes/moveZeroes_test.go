package move_zeroes

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0,1,0,3,12}
	want := []int{1,3,12,0,0}
	MoveZeroes(nums)
	for idx, num := range nums {
		if num != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], num)
		}
	}
}
