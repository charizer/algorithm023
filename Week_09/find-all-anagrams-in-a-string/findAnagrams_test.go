package find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	want := []int{0,6}
	got := FindAnagrams(s, p)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%+v got:%+v", want, got)
	}
}
