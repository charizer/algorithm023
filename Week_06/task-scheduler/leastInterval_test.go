package task_scheduler

import "testing"

func TestLeastInterval(t *testing.T) {
	tasks := []byte{'A','A','A','B','B','B'}
	n := 2
	want := 8
	got := LeastInterval(tasks, n)
	if got != want {
		t.Errorf("want:%d got:%d", want,  got)
	}
}
