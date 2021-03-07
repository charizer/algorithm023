package wildcard_matching

import "testing"

func TestIsMatch(t *testing.T) {
	s := "adceb"
	p := "*a*b"
	want := true
	got := IsMatch(s, p)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
