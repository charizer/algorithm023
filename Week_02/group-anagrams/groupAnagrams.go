package group_anagrams

func GroupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	if len(strs) == 0 {
		return ans
	}
	// 记录每个异位词数据
	m := make(map[[26]byte][]string)
	for _, str := range strs {
		chars := [26]byte{0}
		for i := range str {
			chars[str[i] - 'a']++
		}
		// 为异位词的字符串chars数组值相同，将异位词字符串按chars数组值组合在一起
		m[chars] = append(m[chars], str)
	}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}