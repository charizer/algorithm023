package unique_paths_ii

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) <= 0 || len(obstacleGrid[0]) <= 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	// dp初始化
	for j := 0; j < n; j++ {
		// 如果为障碍物，则dp值为0
		if obstacleGrid[0][j] == 1 {
			dp[j] = 0
		}else{
			// 不是障碍物，第0列为1
			if j == 0 {
				dp[j] = 1
			}else{
				// 不是第0列，则为前一列的值
				dp[j] = dp[j-1]
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			// 障碍物，则dp值设置为0
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			}else{
				// 不是障碍物，且不是第0列，则为上边和下边数值相加
				if j != 0 {
					dp[j] += dp[j-1]
				}
				// 是第0列，则和上边的值相同
			}
		}
	}
	return dp[n-1]
}
