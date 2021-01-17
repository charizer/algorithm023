package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 前序遍历 preorder = [[3],[9],[20,15,7]]
// 中序遍历 inorder = [[9],[3],[15,20,7]]
/*
	3
	/ \
    9  20
	   /  \
	   15   7
 */
func BuildTree(preorder []int, inorder []int) *TreeNode {
	// 递归终止条件
	if len(preorder) == 0 {
		return nil
	}
	// 根节点为前序遍历的第一个节点
	root := &TreeNode{Val: preorder[0]}
	// 根据前序序列中根节点在中序序列中的位置将树分为左、右子树
	splitIdx := 0
	for i := 0 ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			splitIdx = i
			break
		}
	}
	// 递归处理左右子树
	root.Left = BuildTree(preorder[1:splitIdx+1], inorder[:splitIdx])
	root.Right = BuildTree(preorder[splitIdx+1:], inorder[splitIdx+1:])
	return root
}