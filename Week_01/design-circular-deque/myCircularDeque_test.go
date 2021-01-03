package design_circular_deque

import "testing"

func TestConstructor(t *testing.T) {
	deque := Constructor(3)
	want := true
	got := deque.InsertLast(1)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want = true
	got = deque.InsertLast(2)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want = true
	got = deque.InsertFront(3)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want = false
	got = deque.InsertFront(4)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want2 := 2
	got2 := deque.GetRear()
	if got2 != want2 {
		t.Errorf("want:%d got:%d", want2, got2)
	}
	want = true
	got = deque.IsFull()
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want = true
	got = deque.DeleteLast()
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want = true
	got = deque.InsertFront(4)
	if got != want {
		t.Errorf("want:%t got:%t", want, got)
	}
	want2 = 4
	got2 = deque.GetFront()
	if got2 != want2 {
		t.Errorf("want:%d got:%d", want2, got2)
	}
}
