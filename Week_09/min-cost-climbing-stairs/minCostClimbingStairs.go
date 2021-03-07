package min_cost_climbing_stairs

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinCostClimbingStairs(cost []int) int {
	// 小于2阶，直接跳过
	if len(cost) < 2 {
		return 0
	}
	// dp初始化
	f0, f1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		// dp[i] = min(dp[i-1],dp[i-2]) + cost[i]
		f2 := min(f0, f1) + cost[i]
		f0 = f1
		f1 = f2
	}
	// min(dp[len(cost)-1], dp[len(cost)-2])
	return min(f0, f1)
}
