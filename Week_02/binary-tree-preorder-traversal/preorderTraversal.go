package binary_tree_preorder_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归法
func PreorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	// 根节点
	ans = append(ans, root.Val)
	// 左子树
	ans = append(ans, PreorderTraversal(root.Left)...)
	// 右子树
	ans = append(ans, PreorderTraversal(root.Right)...)
	return ans
}

// 迭代法
func PreorderTraversal2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		// 节点值先放入结果，将节点入栈，然后处理左子树
		if p != nil {
			ans = append(ans, p.Val)
			stack = append(stack, p)
			p = p.Left
		}else{
			// 弹出节点，然后处理右子树
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return ans
}
