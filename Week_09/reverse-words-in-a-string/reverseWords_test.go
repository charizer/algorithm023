package reverse_words_in_a_string

import "testing"

func TestReverseWords(t *testing.T) {
	s := "the sky is blue"
	want := "blue is sky the"
	got := ReverseWords(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
	s = "  hello world!  "
	want = "world! hello"
	got = ReverseWords(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}

func TestReverseWords2(t *testing.T) {
	s := "the sky is blue"
	want := "blue is sky the"
	got := ReverseWords2(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
	s = "  hello world!  "
	want = "world! hello"
	got = ReverseWords2(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}
