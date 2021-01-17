package lowest_common_ancestor_of_a_binary_tree

import "testing"

func createTree() (*TreeNode, *TreeNode, *TreeNode) {
	root := &TreeNode{Val: 3}
	l1 := &TreeNode{Val: 5}
	r1 := &TreeNode{Val: 1}
	root.Left = l1
	root.Right = r1
	l1l := &TreeNode{Val: 6}
	l1r := &TreeNode{Val: 2}
	l1.Left = l1l
	l1.Right = l1r
	l2l := &TreeNode{Val: 7}
	l2r := &TreeNode{Val: 4}
	l1r.Left = l2l
	l1r.Right = l2r
	r1l := &TreeNode{Val: 0}
	r1r := &TreeNode{Val: 8}
	r1.Left = r1l
	r1.Right = r1r
	return root, l1, r1
}

func TestLowestCommonAncestor(t *testing.T) {
	root, p , q := createTree()
	want := 3
	got := LowestCommonAncestor(root, p, q)
	if got.Val != want {
		t.Errorf("want:%d got:%d", 3, got.Val)
	}
}

