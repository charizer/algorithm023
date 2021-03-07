package decode_ways

func NumDecodings(s string) int {
	if len(s) <= 0 || s[0] == '0' {
		return 0
	}
	n := len(s)
	//dp[n]代表以s[n-1]为结尾的解码数
	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < n; i++ {
		// 当前为'0'，则前一个值需要时'1'或'2'
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				// 此种情况s[i]s[i-1]被唯一编码，不增加情况
				dp[i+1] = dp[i-1]
			}else{
				// s[i-1]不是'1'或'2'，则不能编码
				return 0
			}
		}else{
			// 当前位和前一位要小于26，则当前为和前一位可以单独编码也可以连起来编码
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
				// 分开编码为dp[i], 合并编码为dp[i-1]
				dp[i+1] = dp[i] + dp[i-1]
			}else{
				// 否则只能分开编码，则为dp[i]
				dp[i+1] = dp[i]
			}
		}
	}
	return dp[n]
}
