package combinations

func Combine(n int, k int) [][]int {
	ans := [][]int{}
	if n <= 0 || k <= 0 {
		return ans
	}
	backtrack(1, n, k, []int{}, &ans)
	return ans
}

// 回溯解法框架
func backtrack(idx, n, k int, path []int, ans *[][]int){
	// 递归终止条件，已经选出k个数
	if len(path) == k {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := idx; i <= n; i++ {
		// 选择i
		path = append(path, i)
		// 选其他数
		backtrack(i + 1, n, k, path, ans)
		// 不选i
		path = path[0:len(path)-1]
	}
}
