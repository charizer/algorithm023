package maximal_square

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// 边长
	ans := 0
	// 二维dp数组
	m, n := len(matrix),len(matrix[0])
	dp := make([][]int, m)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}
	// dp方程
	// dp[row][col] = min(dp[row-1][col],min(dp[row][col-1], dp[row-1][col-1])) + 1
	for row := range matrix {
		for col := range matrix[row] {
			// 当前cell值为'0'时，则dp值则为0
			if matrix[row][col] == '0' {
				dp[row][col] = 0
			} else if row == 0 || col == 0 {
				// 在第0行或第0列，dp值为1，为'0'的情况已经在前边处理了
				dp[row][col] = 1
			}else{
				// 其他情况直接应用dp方程
				dp[row][col] = min(dp[row-1][col],min(dp[row][col-1], dp[row-1][col-1])) + 1
			}
			// 更新最大边长
			if dp[row][col] > ans {
				ans =dp[row][col]
			}
		}
	}
	return ans * ans
}
