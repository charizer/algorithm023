package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 到叶子节点或者p、q中任意节点，递归终止
	if root == nil || root == p || root == q {
		return root
	}
	// 左子树中查找p、q
	left := LowestCommonAncestor(root.Left, p, q)
	// 右子树中查找p、q
	right := LowestCommonAncestor(root.Right, p, q)
	// p、q分别在左右子树中，则当前节点即为所求
	if left != nil && right != nil {
		return root
	}
	// 在子树的一边，则非nil的left或right为所求
	if left == nil {
		return right
	}
	return left
}