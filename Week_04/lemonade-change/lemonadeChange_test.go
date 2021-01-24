package lemonade_change

import "testing"

func TestLemonadeChange(t *testing.T) {
	bills := []int{5,5,5,10,20}
	got := LemonadeChange(bills)
	want := true
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	bills = []int{5,5,10}
	got = LemonadeChange(bills)
	want = true
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	bills = []int{10,10}
	got = LemonadeChange(bills)
	want = false
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	bills = []int{5,5,10,10,20}
	got = LemonadeChange(bills)
	want = false
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	bills = []int{}
	got = LemonadeChange(bills)
	want = true
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
