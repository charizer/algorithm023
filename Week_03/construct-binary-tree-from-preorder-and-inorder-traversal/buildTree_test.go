package construct_binary_tree_from_preorder_and_inorder_traversal

import "testing"

func preOrderTravel(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, preOrderTravel(root.Left)...)
	ans = append(ans, preOrderTravel(root.Right)...)
	return ans
}

func inOrderTravel(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, inOrderTravel(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inOrderTravel(root.Right)...)
	return ans
}

func TestBuildTree(t *testing.T) {
	preOrder := []int{3,9,20,15,7}
	inOrder := []int{9,3,15,20,7}
	got := BuildTree(preOrder, inOrder)
	preGot := preOrderTravel(got)
	inGot := inOrderTravel(got)
	for idx := range preOrder {
		if preGot[idx] != preOrder[idx] {
			t.Errorf("pre want:%d got:%d", preOrder[idx], preGot[idx])
		}
	}
	for idx := range inOrder {
		if inGot[idx] != inOrder[idx] {
			t.Errorf("in want:%d got:%d", inOrder[idx], inGot[idx])
		}
	}
}
