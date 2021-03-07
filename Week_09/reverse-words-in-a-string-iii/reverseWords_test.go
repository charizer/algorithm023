package reverse_words_in_a_string_iii

import "testing"

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	want := "s'teL ekat edoCteeL tsetnoc"
	got := ReverseWords(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}
