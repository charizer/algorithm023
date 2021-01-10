package group_anagrams

import "testing"

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	want := [][]string{{"eat","tea","ate"},{"tan", "nat"}, {"bat"} }
	got := GroupAnagrams(strs)
	for out := range got {
		for inner := range got[out] {
			if got[out][inner] != want[out][inner] {
				t.Errorf("want:%s got:%s", want[out][inner], got[out][inner])
			}
		}
	}
}
