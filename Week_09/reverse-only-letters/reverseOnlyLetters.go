package reverse_only_letters

func ReverseOnlyLetters(S string) string {
	bs := []byte(S)
	for left, right := 0, len(bs) - 1; left < right; {
		// 找到左边的字母
		for left < right && !isletter(bs[left]) {
			left++
		}
		// 找到右边的字母
		for left < right && !isletter(bs[right]) {
			right--
		}
		// 进行交换
		bs[left], bs[right] = bs[right], bs[left]
		// 继续处理后续字母
		left++
		right--
	}
	return string(bs)
}

// 判断是否是字母
func isletter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}
