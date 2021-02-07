package student_attendance_record_ii

import "testing"

func TestCheckRecord(t *testing.T) {
	n := 2
	want := 8
	got := CheckRecord(n)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
