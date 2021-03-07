package longest_valid_parentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	want := 2
	got := LongestValidParentheses(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = ")()())"
	want = 4
	got = LongestValidParentheses(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestLongestValidParentheses2(t *testing.T) {
	s := "(()"
	want := 2
	got := LongestValidParentheses2(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = ")()())"
	want = 4
	got = LongestValidParentheses2(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
