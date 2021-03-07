package race_car

func Racecar(target int) int {
	dp := make([]int, target + 2)
	dp[1] = 1 //A
	dp[2] = 4 //AARA 或者 ARRA
	k := 2
	// S记录连续k个A指令，达到的位置
	S := 3
	for i := 3; i <= target; i++ {
		if i == S {
			dp[i] = k
			k++
			// 2^k - 1
			S = (1<<k) - 1
		} else {
			// 情况1：连续k个A后，回退
			dp[i] = k + 1 + dp[S-i]
			// 情况2：连续k-1个A后，回退(0/1/.../k-2)步后，再前进
			for back := 0; back <= k-2; back++ {
				// 回退后还需前进的距离：i+S(back)-S(k-1)
				distance := i + (1<<back) - (1<<(k-1))
				dp[i] = min(dp[i], (k-1)+2+back+dp[distance])
			}
		}
	}
	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
