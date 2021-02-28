package relative_sort_array

import (
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T){
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	want := []int{2,2,2,1,4,3,3,9,6,7,19}
	got := RelativeSortArray(arr1, arr2)
	if !reflect.DeepEqual(want, got){
		t.Errorf("want:%+v got:%+v", want, got)
	}
}

func TestRelativeSortArray2(t *testing.T){
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	want := []int{2,2,2,1,4,3,3,9,6,7,19}
	got := RelativeSortArray2(arr1, arr2)
	if !reflect.DeepEqual(want, got){
		t.Errorf("want:%+v got:%+v", want, got)
	}
}
