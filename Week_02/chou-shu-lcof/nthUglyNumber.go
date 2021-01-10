package chou_shu_lcof

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}else{
			return c
		}
	}else{
		if b < c {
			return b
		}else{
			return c
		}
	}
}

func NthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	// 用于存放n个丑数
	dp := make([]int, n)
	// 记录不同丑数因子的丑数指针，初始化都指向第一个丑数
	two, three, five := 0, 0, 0
	// 第一个丑数
	dp[0] = 1
	// 生成剩余的n-1个丑数
	for i := 1; i < n; i++ {
		// 获取下一个最小的丑数
		next := min(dp[two]*2, dp[three]*3, dp[five]*5)
		dp[i] = next
		// 更新当前2、3、5丑数指针的位置
		/*
		// 2、3、5倍数会有重复的，一次不能只更新一个，而必须全部更新
		switch next {
		case dp[two]*2:
			two++
		case dp[three]*3:
			three++
		case dp[five]*5:
			five++
		}*/
		if next == dp[two] * 2 {
			two++
		}
		if next == dp[three] * 3 {
			three++
		}
		if next == dp[five] * 5 {
			five++
		}
	}
	return dp[n-1]
}
