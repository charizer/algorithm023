package find_all_anagrams_in_a_string

func FindAnagrams(s string, p string) []int {
	ans := []int{}
	// win: 窗口字符串的统计， pattern：p字符串的统计
	win, pattern := make(map[byte]int), make(map[byte]int)
	for i := range p {
		pattern[p[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s) {
		ch := s[right]
		// 字符是p中字符
		if _, ok := pattern[ch]; ok {
			win[ch]++
			// 个数相同，则有效字符加1
			if win[ch] == pattern[ch] {
				valid++
			}
		}
		// 窗口长度达到p的长度，需要收缩
		if right - left + 1 >= len(p) {
			// 完全匹配，记录结果
			if valid == len(pattern) {
				ans = append(ans, left)
			}
			// 将左边字符移除窗口
			ch := s[left]
			left++
			// 移除的字符是p中的一个字符，需要进一步更新匹配记录
			if _, ok := pattern[ch]; ok {
				if win[ch] == pattern[ch] {
					valid--
				}
				win[ch]--
			}
		}
		right++
	}
	return ans
}
