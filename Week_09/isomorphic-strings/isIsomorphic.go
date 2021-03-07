package isomorphic_strings

func IsIsomorphic(s string, t string) bool {
	// 长度不同，肯定不是同构字符串
	if len(s) != len(t) {
		return false
	}
	// 存放s对应的t
	s2t := make(map[byte]byte)
	// 存放t对应的s
	t2s := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		// 记录两个map的对应关系
		if _, ok := s2t[s[i]]; !ok {
			s2t[s[i]] = t[i]
		}
		if _, ok := t2s[t[i]]; !ok {
			t2s[t[i]] = s[i]
		}
		// 不匹配则直接返回
		if s2t[s[i]] != t[i] || t2s[t[i]] != s[i] {
			return false
		}
	}
	return true
}
