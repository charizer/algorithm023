package reverse_only_letters

import "testing"

func TestReverseOnlyLetters(t *testing.T) {
	s := "a-bC-dEf-ghIj"
	want := "j-Ih-gfE-dCba"
	got := ReverseOnlyLetters(s)
	if got != want {
		t.Errorf("want:%s got:%s", want, got)
	}
}
