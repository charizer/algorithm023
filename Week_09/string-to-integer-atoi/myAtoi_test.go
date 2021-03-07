package string_to_integer_atoi

import "testing"

func TestMyAtoi(t *testing.T) {
	s := "42"
	want := 42
	got := MyAtoi(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "   -42"
	want = -42
	got = MyAtoi(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "4193 with words"
	want = 4193
	got = MyAtoi(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "words and 987"
	want = 0
	got = MyAtoi(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "-91283472332"
	want = -2147483648
	got = MyAtoi(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "91283472332"
	want = 1 << 31 - 1
	got = MyAtoi(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	s = "-2147483649"
	want = -1 << 31
	got = MyAtoi(s)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
