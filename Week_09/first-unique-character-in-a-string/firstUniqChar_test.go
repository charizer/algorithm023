package first_unique_character_in_a_string

import "testing"

func TestFirstUniqChar(t *testing.T) {
	s := "leetcode"
	want := 0
	got := FirstUniqChar(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "loveleetcode"
	want = 2
	got = FirstUniqChar(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
