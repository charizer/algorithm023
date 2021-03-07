package distinct_subsequences

import "testing"

func TestNumDistinct(t *testing.T) {
	s := "rabbbit"
	ts := "rabbit"
	want := 3
	got := NumDistinct(s, ts)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
