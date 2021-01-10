package binary_tree_preorder_traversal

import "testing"

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{Val:1}
	right1 := &TreeNode{Val: 2}
	left2 := &TreeNode{Val: 3}
	root.Right = right1
	right1.Left = left2
	want := []int{1,2,3}
	got := PreorderTraversal(root)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d g:%d", want[idx], g)
		}
	}
}

func TestPreorderTraversal2(t *testing.T) {
	root := &TreeNode{Val:1}
	right1 := &TreeNode{Val: 2}
	left2 := &TreeNode{Val: 3}
	root.Right = right1
	right1.Left = left2
	want := []int{1,2,3}
	got := PreorderTraversal2(root)
	for idx, g := range got {
		if g != want[idx] {
			t.Errorf("want:%d g:%d", want[idx], g)
		}
	}
}
