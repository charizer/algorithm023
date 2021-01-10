package n_ary_tree_preorder_traversal

type Node struct {
	Val int
	Children []*Node
}

// 递归解法
func Preorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	// 先遍历根节点
	ans = append(ans, root.Val)
	// 然后递归遍历各孩子节点
	for _, child := range root.Children {
		ans = append(ans, Preorder(child)...)
	}
	return ans
}

// 迭代解法
func Preorder2(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*Node{root}
	// 栈中还有节点，继续处理
	for len(stack) > 0 {
		// 弹出节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		// 结果是先左后右，所以孩子从右到左入栈
		for i := len(node.Children) -1; i >=0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ans
}


