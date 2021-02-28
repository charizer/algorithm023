package merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	want := [][]int{{1,6},{8,10},{15,18}}
	got := Merge(intervals)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%+v got:%+v", want, got)
	}
	intervals = [][]int{{1,4},{4,5}}
	want = [][]int{{1,5}}
	got = Merge(intervals)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%+v got:%+v", want, got)
	}
}
