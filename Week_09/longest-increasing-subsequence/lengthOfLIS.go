package longest_increasing_subsequence

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	// dp数组初始化，dp[i]代表以nums[i]结尾的最长子序列长度
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// 前面的数字比当前数字小，把可以把当前数字拼到结尾形成一个子序列，最终值取较大值
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	// 选取最大递增子序列长度
	ans := 0
	for i := range dp {
		ans = max(ans, dp[i])
	}
	return ans
}
