package plus_one

func PlusOne(digits []int) []int {
	if len(digits) <= 0 {
		return []int{}
	}
	// 用于记录进位, 初始为1，当做加1
	plus := 1
	for idx := len(digits)-1; idx >= 0; idx-- {
		// 没有进位了，直接返回digits即为所求
		if plus == 0 {
			return digits
		}
		// 当前位加上进位
		digits[idx] = digits[idx] + plus
		// 可能会产生进位，进行处理
		plus, digits[idx] = digits[idx] / 10, digits[idx] % 10
	}
	// 有向最高位的进位，进行处理
	if plus == 1 {
		digits = append([]int{plus}, digits...)
	}
	return digits
}
