package frog_jump

func CanCross(stones []int) bool {
	if len(stones) <= 0 {
		return false
	}
	dp := make(map[int]map[int]struct{}, len(stones))
	// 初始化每块石头及跳到的次数
	for _, i := range stones {
		dp[i] = make(map[int]struct{})
	}
	// 第0块石头初始化跳数0
	dp[0][0] = struct{}{}
	for i := 0; i < len(stones); i++ {
		for k, _ := range dp[stones[i]] {
			// 从当前石头跳k-1~k+1到另一块石头
			for step := k - 1; step <= k + 1; step++ {
				if step > 0 {
					if _, ok := dp[stones[i] + step]; ok {
						// 能够跳到另一块石头，记录下跳的步数step
						dp[stones[i] + step][step] = struct{}{}
					}
				}
			}
		}
	}
	return len(dp[stones[len(stones)-1]]) != 0
}
