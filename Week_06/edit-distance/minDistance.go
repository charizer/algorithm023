package edit_distance

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
	增，dp[i][j] = dp[i][j - 1] + 1
	删，dp[i][j] = dp[i - 1][j] + 1
	改，dp[i][j] = dp[i - 1][j - 1] + 1
	状态转移方程：
		dp[i][j] = min(dp[i - 1][j] ， dp[i][j - 1] ， dp[i - 1][j - 1]) + 1
 	如果刚好这两个字母相同 word1[i - 1] = word2[j - 1] ，那么可以直接参考 dp[i - 1][j - 1] ，操作不用加一
*/

func MinDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	if n1 == 0 && n2 == 0 {
		return 0
	}
	dp := make([][]int, n1 + 1)
	for  i := range dp {
		dp[i] = make([]int, n2 + 1)
	}
	dp[0][0] = 0
	// 第0列
	for i := 1; i <= n2; i++ {
		dp[0][i] = i
	}
	// 第0行
	for i := 1; i <= n1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n1; i ++ {
		for j := 1; j <= n2; j++ {
			dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			// 两个字符刚好一样，那么可以直接参考 dp[i - 1][j - 1]
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j],dp[i-1][j-1])
			}
		}
	}
	return dp[n1][n2]
}
