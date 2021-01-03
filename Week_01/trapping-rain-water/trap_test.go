package trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	want := 6
	got := Trap(height)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	height = []int{4,2,0,3,2,5}
	want = 9
	got = Trap(height)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
