package word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	beginWord, endWord := "hit", "cog"
	wordList := []string{"hot","dot","dog","lot","log","cog"}
	want := 5
	got := LadderLength(beginWord, endWord, wordList)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
