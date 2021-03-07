package race_car

import "testing"

func TestRacecar(t *testing.T) {
	target := 6
	want := 5
	got := Racecar(target)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
