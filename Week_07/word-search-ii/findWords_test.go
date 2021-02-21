package word_search_ii

import "testing"

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	want := []string{"oath", "eat"}
	got := FindWords(board, words)
	for idx := range want {
		if got[idx] != want[idx] {
			t.Errorf("want:%s got:%s", want[idx], got[idx])
		}
	}

	board = [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words = []string{"oath", "pea", "eat", "rain", "oathi", "oathk", "oathf", "oate", "oathii", "oathfi", "oathfii"}
	want = []string{"oath", "oathk", "oathf", "oathfi", "oathfii", "oathi", "oathii", "oate", "eat"}
	got = FindWords(board, words)
	/*for idx := range want {
		if got[idx] != want[idx] {
			t.Errorf("want:%s got:%s", want[idx], got[idx])
		}
	}*/
	t.Logf("want:%+v got:%+v", want, got)
}
