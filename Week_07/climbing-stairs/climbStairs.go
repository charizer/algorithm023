package climbing_stairs

func ClimbStairs(n int) int {
	if n < 2 {
		return 1
	}
	f0, f1, f2 := 1, 1, 2
	for i := 3; i <= n; i++ {
		f0 = f1
		f1 = f2
		f2 = f0 + f1
	}
	return f2
}
