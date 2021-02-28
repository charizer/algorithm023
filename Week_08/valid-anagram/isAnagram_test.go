package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	s2 := "nagaram"
	want := true
	got := IsAnagram(s, s2)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}