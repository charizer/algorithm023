package palindromic_substrings

/*
dp状态转移分析：
当s[i]与s[j]相等时,有如下三种情况
情况一：下标i 与 j相同，同一个字符例如a，是回文子串
情况二：下标i 与 j相差为1，例如aa，也是文子串
情况三：下标：i 与 j相差大于1的时候，此时s[i]与s[j]已经相同了，看dp[i + 1][j - 1]是否为true。
情况一和情况二统一为：j - i < 2
*/

func CountSubstrings(s string) int {
	n, count := len(s), 0
	dp := make([][]bool,n)
	for idx := range dp {
		dp[idx] = make([]bool, n)
	}
	for i := n -1; i >= 0; i-- {
		for j := i;j < n; j++ {
			if s[i] == s[j] && (j -i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				count++
			}
		}
	}
	return count
}
