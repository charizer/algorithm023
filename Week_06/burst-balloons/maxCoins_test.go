package burst_balloons

import "testing"

func TestMaxCoins(t *testing.T) {
	nums := []int{3,1,5,8}
	want := 167
	got := MaxCoins(nums)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}