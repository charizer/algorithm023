package wildcard_matching

func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}
	// 两个空串，匹配
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		// s为空串，则p每个字符均是*，则匹配
		dp[0][j] = p[j-1] == '*' && dp[0][j-1]
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 当前字符相等或p当前字符是 '?' 可以匹配任何一个非''，则dp[i][j]可以从dp[i-1][j-1]转移过来
			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else if p[j-1] == '*' {
				// *匹配'' 或者 * 匹配多个字符
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}

		}
	}
	return dp[m][n]
}
