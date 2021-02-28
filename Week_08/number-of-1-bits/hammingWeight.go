package number_of_1_bits

func HammingWeight(num uint32) int {
	ans := 0
	// 利用位运算性质，num & (num -1) 能把最后一位1消掉
	for num != 0 {
		ans++
		num = num & (num - 1)
	}
	return ans
}
