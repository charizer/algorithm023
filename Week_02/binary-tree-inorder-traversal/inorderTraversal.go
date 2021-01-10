package binary_tree_inorder_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归法
func InorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	// 左子树
	ans = append(ans, InorderTraversal(root.Left)...)
	// 根节点
	ans = append(ans, root.Val)
	// 右子树
	ans = append(ans, InorderTraversal(root.Right)...)
	return ans
}

// 迭代法
func InorderTraversal2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		// 节点非空，则入栈，继续处理左孩子
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		}else{
			// 节点是空，则弹出栈顶，栈顶值记录进结果，然后处理右孩子
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, p.Val)
			p = p.Right
		}
	}
	return ans
}
