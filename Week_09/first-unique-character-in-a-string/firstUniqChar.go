package first_unique_character_in_a_string

func FirstUniqChar(s string) int {
	ms := make(map[byte]int)
	// 统计各字符次数
	for i := range s {
		ms[s[i]]++
	}
	// 再次遍历字符串，找到第一个唯一字符返回其位置
	for i := range s {
		if ms[s[i]] == 1 {
			return i
		}
	}
	return -1
}
