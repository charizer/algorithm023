package chou_shu_lcof

import "testing"

func TestNthUglyNumber(t *testing.T) {
	n := 10
	want := 12
	got := NthUglyNumber(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
