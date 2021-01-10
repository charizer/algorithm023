package n_ary_tree_level_order_traversal

type Node struct {
	Val int
	Children []*Node
}

// dfs解法
func LevelOrder(root *Node) [][]int {
	ans := [][]int{}
	dfs(&ans, root, 0)
	return ans
}

func dfs(ans *[][]int, root *Node, level int) {
	if root == nil {
		return
	}
	// 需要建立新层
	if level >= len(*ans) {
		*ans = append(*ans, []int{})
	}
	// 将节点值放入当前层
	(*ans)[level] = append((*ans)[level], root.Val)
	// 继续处理孩子节点
	for _, child := range root.Children {
		dfs(ans, child, level+1)
	}
}

// bfs解法
func LevelOrder2(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		// 存放当前层结果
		sub := []int{}
		// 当前层节点个数
		size := len(queue)
		for i := 0; i < size; i++ {
			// 将节点出队列并记录进当前层结果
			node := queue[0]
			queue = queue[1:]
			sub = append(sub, node.Val)
			// 将节点孩子入队列，用于下一层遍历
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		// 一层处理完，将当前层结果放入结果集中
		ans = append(ans, sub)
	}
	return ans
}