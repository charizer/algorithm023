package valid_anagram

func IsAnagram(s string, t string) bool {
	// 如果长度不等，肯定不是字母异位词
	if len(s) != len(t) {
		return false
	}
	// 记录各字符次数
	m := make(map[rune]int)
	// 遍历s时各次数+1
	for _, ch := range s {
		m[ch]++
	}
	// 遍历t时各字符次数-1，如果次数<0, 则说明s的次数小于t中次数，可以提前结束
	for _, ch := range t {
		m[ch]--
		if m[ch] < 0 {
			return false
		}
	}
	/* 不需要，最开始先判断了两个字符串长度相等与否，如果还有不为0的，肯定在遍历t时已经提前结束了
	// 如果遍历完两个字符串，还有次数不是0的，说明肯定不是字母异位词
	for _, v := range m {
		if v != 0 {
			return false
		}
	}*/
	return true
}
