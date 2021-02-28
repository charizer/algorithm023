package valid_anagram

func IsAnagram(s string, t string) bool {
	// 长度不等，肯定不是异位词
	if len(s) != len(t) {
		return false
	}
	// 使用map统计各字符数量
	hash := make(map[byte]int)
	// s中出现，则相应字符+1，t中出现，则相应字符-1
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
		hash[t[i]]--
	}
	// 经过上述处理，如果所有字符个数都是0，则才是有效的字母异位词
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
