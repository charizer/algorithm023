package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	want := true
	got := IsAnagram(s1,s2)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	s1 = "rat"
	s2 = "car"
	want = false
	got = IsAnagram(s1,s2)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
