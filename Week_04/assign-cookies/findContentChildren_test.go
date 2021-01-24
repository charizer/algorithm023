package assign_cookies

import "testing"

func TestFindContentChildren(t *testing.T) {
	g := []int{1,2,3}
	s := []int{1, 1}
	want := 1
	got := FindContentChildren(g, s)
	if got != want {
		t.Errorf("want:%d got:%d", want,  got)
	}
	g = []int{1,2}
	s = []int{1, 2, 3}
	want = 2
	got = FindContentChildren(g, s)
	if got != want {
		t.Errorf("want:%d got:%d", want,  got)
	}
	g = []int{10, 9 ,8 , 7}
	s = []int{5, 6, 7, 8}
	want = 2
	got = FindContentChildren(g, s)
	if got != want {
		t.Errorf("want:%d got:%d", want,  got)
	}
}