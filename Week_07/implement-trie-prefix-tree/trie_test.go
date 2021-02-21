package implement_trie_prefix_tree

import "testing"

func TestConstructor(t *testing.T) {
	tree := Constructor()
	tree.Insert("apple")
	want := true
	got := tree.Search("apple")
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want = false
	got = tree.Search("app")
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want = true
	got = tree.StartsWith("app")
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	tree.Insert("app")
	want = true
	got = tree.Search("app")
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
}
