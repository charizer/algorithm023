package power_of_two

func IsPowerOfTwo(n int) bool {
	// 二进制中只有一个1，则说明是2的幂
	return n > 0 && n & (n-1) == 0
}
