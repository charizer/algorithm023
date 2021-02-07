package minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	ts := "ABC"
	want := "BANC"
	got := MinWindow(s, ts)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}
