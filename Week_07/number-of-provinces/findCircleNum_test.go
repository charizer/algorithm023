package number_of_provinces

import "testing"

func TestFindCircleNum(t *testing.T) {
	isConnected := [][]int{{1,1,0},{1,1,0},{0,0,1}}
	want := 2
	got := FindCircleNum(isConnected)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
