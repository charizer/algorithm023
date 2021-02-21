package generate_parentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	want := []string{"((()))","(()())","(())()","()(())","()()()"}
	got := GenerateParenthesis(n)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want:%+v, got:%+v", want, got)
	}
}