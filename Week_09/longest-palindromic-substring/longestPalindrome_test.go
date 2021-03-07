package longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	want := "bab"
	got := LongestPalindrome(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
	s = "cbbd"
	want = "bb"
	got = LongestPalindrome(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}
