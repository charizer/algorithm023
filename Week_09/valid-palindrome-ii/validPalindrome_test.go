package valid_palindrome_ii

import "testing"

func TestValidPalindrome(t *testing.T) {
	s := "aba"
	want := true
	got := ValidPalindrome(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s = "abca"
	want = true
	got = ValidPalindrome(s)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
