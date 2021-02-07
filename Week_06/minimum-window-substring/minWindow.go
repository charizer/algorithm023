package minimum_window_substring

import "math"

func MinWindow(s string, t string) string {
	ans, cnt := "", math.MaxInt32
	// 用于记录字符出现次数
	hm := make(map[byte]int)
	for idx := range t {
		hm[t[idx]]++
	}
	l, r := 0, 0
	for ; r < len(s); r++ {
		hm[s[r]]--
		// s中已经包含全部t，则缩小窗口，找最小子串
		for check(hm) {
			if r - l + 1 < cnt {
				cnt = r -l + 1
				ans = s[l: r+1]
			}
			hm[s[l]]++
			l++
		}
	}
	return ans
}

// 检查s中是否完全包含t
func check(hm map[byte]int) bool {
	for _, v := range hm {
		if v > 0 {
			return false
		}
	}
	return true
}

