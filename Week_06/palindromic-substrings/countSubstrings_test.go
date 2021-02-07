package palindromic_substrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	s := "abc"
	want := 3
	got := CountSubstrings(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "aaa"
	want = 6
	got = CountSubstrings(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
