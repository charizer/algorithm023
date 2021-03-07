package longest_palindromic_substring

func LongestPalindrome(s string) string {
	if len(s) <= 0 {
		return s
	}
	low, high := 0, 0
	for i := 0; i < len(s); i++ {
		//寻找长度为奇数的回文串
		l, r := palindrome(s, i, i)
		// 区间长度更大，则更新
		if r - l > high - low {
			low, high = l, r
		}
		//寻找长度为偶数的回文串
		l, r = palindrome(s, i, i+1)
		// 区间长度更大，则更新
		if r - l > high - low {
			low, high = l, r
		}
	}
	// 最长回文
	return s[low:high+1]
}

// 从l,r向两端扩散，返回最长回文区间
func palindrome(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
