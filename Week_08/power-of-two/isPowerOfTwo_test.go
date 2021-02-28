package power_of_two

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	n := 16
	want := true
	got := IsPowerOfTwo(n)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
