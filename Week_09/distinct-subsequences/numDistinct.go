package distinct_subsequences

func NumDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	// 两个都是空字符串
	dp[0][0] = 1
	// t为空串
	for i := 1; i <= m; i++ {
		dp[i][0] = 1
	}
	// s为空串
	for j := 1; j <= n; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				// dp[i-1][j-1]: 用s[i]来匹配
				// dp[i-1][j]: 不用si[i]来匹配
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}else{
				// 不用si[i]来匹配
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
