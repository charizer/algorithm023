package reverse_words_in_a_string_iii

func ReverseWords(s string) string {
	bs := []byte(s)
	for left := 0; left < len(bs);{
		right := left
		// 右指针右移到单词结尾
		for right < len(bs) && bs[right] != ' ' {
			right++
		}
		// right的后一个字符是下一个单词的开始
		pos := right + 1
		// right的前一个字符是单词结尾
		end := right-1
		// 反转当前单词
		for left < end {
			bs[left],bs[end] = bs[end],bs[left]
			left++
			end--
		}
		// 左指针移到下一个单词开始
		left = pos
	}
	return string(bs)
}
