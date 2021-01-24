package word_ladder_ii

import "testing"

func TestFindLadders(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot","dot","dog","lot","log","cog"}
	want := [][]string{
		{"hit","hot","dot","dog","cog"},
		{"hit","hot","lot","log","cog"},
	}
	got := FindLadders(beginWord, endWord, wordList)
	for idx := range want {
		for i := range want[idx] {
			if got[idx][i] != want[idx][i] {
				t.Errorf("want:%s got:%s", want[idx][i], got[idx][i])
			}
		}
	}
}