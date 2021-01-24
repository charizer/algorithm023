package word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList:= []string{"hot","dot","dog","lot","log","cog"}
	want := 5
	got := LadderLength(beginWord, endWord, wordList)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot","dot","dog","lot","log"}
	want = 0
	got = LadderLength(beginWord, endWord, wordList)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}

func TestLadderLength2(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList:= []string{"hot","dot","dog","lot","log","cog"}
	want := 5
	got := LadderLength2(beginWord, endWord, wordList)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot","dot","dog","lot","log"}
	want = 0
	got = LadderLength2(beginWord, endWord, wordList)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
