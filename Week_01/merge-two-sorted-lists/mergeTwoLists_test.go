package merge_two_sorted_lists

import "testing"

func createList(arr []int) *ListNode {
	head := &ListNode{}
	pre := head
	for _, item := range arr {
		pre.Next = &ListNode{Val: item}
		pre = pre.Next
	}
	return head.Next
}

func TestMergeTwoLists(t *testing.T) {
	l1 := createList([]int{1,2,4})
	l2 := createList([]int{1,3,4})
	got := MergeTwoLists(l1, l2)
	want := []int{1,1,2,3,4,4}
	idx := 0
	for got != nil {
		if got.Val != want[idx] {
			t.Errorf("want:%d got:%d", want[idx], got.Val)
			break
		}
		idx++
		got = got.Next
	}
}
