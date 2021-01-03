package merge_two_sorted_lists

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 记录新链表头节点，为哑节点， newHead.Next为实际链表
	newHead := &ListNode{}
	// 用于遍历用
	pre := newHead
	// 只要有一个链表没遍历完就继续处理
	for l1 != nil || l2 != nil {
		// 如果两个链表都没有遍历完，则比较大小，较小链表为所求，并向后移动
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				pre.Next = l1
				l1 = l1.Next
				pre = pre.Next
			}else{
				pre.Next = l2
				l2 = l2.Next
				pre = pre.Next
			}
			// 已经有一个链表遍历完成，只需将未遍历完的拼接在新链表的尾部就可以结束了
		}else if l1 != nil {
			pre.Next = l1
			break
		}else if l2 != nil {
			pre.Next = l2
			break
		}
	}
	return newHead.Next
}
