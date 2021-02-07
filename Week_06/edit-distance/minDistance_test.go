package edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	want := 3
	got := MinDistance(word1, word2)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
