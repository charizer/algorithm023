package isomorphic_strings

import "testing"

func TestIsIsomorphic(t *testing.T) {
	s := "egg"
	ts := "add"
	want := true
	got := IsIsomorphic(s, ts)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "foo"
	ts = "bar"
	want = false
	got = IsIsomorphic(s, ts)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "badc"
	ts = "baba"
	want = false
	got = IsIsomorphic(s, ts)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
