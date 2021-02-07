package burst_balloons

// 动态转移方程：dp[i][j] = dp[i][k] + dp[k][j]
func MaxCoins(nums []int) int {
	n := len(nums)
	ans := make([][]int, n + 2)
	for i := 0; i < n + 2; i++ {
		ans[i] = make([]int, n + 2)
	}
	// 前后增加一个元素，简化处理
	val := make([]int, n + 2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n + 1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += ans[i][k] + ans[k][j]
				ans[i][j] = max(ans[i][j], sum)
			}
		}
	}
	return ans[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
