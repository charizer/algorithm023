package reverse_string_ii

import "testing"

func TestReverseStr(t *testing.T) {
	s := "abcdefg"
	k := 2
	want := "bacdfeg"
	got := ReverseStr(s, k)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}
