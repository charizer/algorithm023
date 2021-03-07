package decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	s := "12"
	want := 2
	got := NumDecodings(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "226"
	want = 3
	got = NumDecodings(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
