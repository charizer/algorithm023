package minimum_path_sum

func MinPathSum(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	// 行、列
	m, n := len(grid), len(grid[0])
	// dp问题， dp公式为: dp[i][j] = min(dp[i][j-1],dp[i-1][j]) + grid[i][j]
	// 可通过一维数组简化
	dp := make([]int, n)
	dp[0] = grid[0][0]
	// dp数组初始化, 初始化第0行的处理
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	// 第1行开始
	for i := 1; i < m; i++ {
		// 第0列为上一行此列值加上当前表格的值
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			// 当前值为上一行列值或当前行前一列值中较小值加上当前表格的值
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
