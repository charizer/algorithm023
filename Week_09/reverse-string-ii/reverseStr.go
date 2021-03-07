package reverse_string_ii

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ReverseStr(s string, k int) string {
	bs := []byte(s)
	// 每次增加2k，处理下一个2k长度字符串
	for i := 0; i < len(bs); i += 2*k {
		// 左位置
		left := i
		// k个字符或当前位置到字符串结尾的较小值，为要反转的子字符串
		right := min(i + k - 1, len(bs) - 1)
		for left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		}
	}
	return string(bs)
}
