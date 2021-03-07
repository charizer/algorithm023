package valid_palindrome_ii

func ValidPalindrome(s string) bool {
	left, right := 0, len(s) -1
	for left < right {
		// 如果不相等，去掉左边或右边的一个字符再判断剩余字符串是否回文
		if s[left] != s[right] {
			return validPalindrome(s, left+1, right) || validPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

// 判断指定左右区间的字符是否回文
func validPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
