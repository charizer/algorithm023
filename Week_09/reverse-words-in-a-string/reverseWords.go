package reverse_words_in_a_string

import "strings"

func ReverseWords(s string) string {
	ans := ""
	bs := []byte(s)
	// 整个字符串先反转一次
	left, right := 0, len(bs)-1
	for left < right {
		bs[left], bs[right] = bs[right], bs[left]
		left++
		right--
	}
	// 然后再逐个反转单词
	for left = 0; left < len(bs); {
		// 左指针跳过空格
		for left < len(bs) && bs[left] == ' ' {
			left++
		}
		// 右指针一直移动直到遇到新空格
		for right = left; right < len(bs) && bs[right] != ' '; right++ {
		}
		// 左右指针相等，说明没有非空格的字符了，跳出循环
		if right == left {
			break
		}
		// 左右指针之间为一个单词，从右到左记录下单词的字符，即为当前反转的单词
		for i := right-1; i >= left; i-- {
			ans += string(bs[i])
		}
		// 加上一个空格，处理下一个单词
		ans += " "
		left = right
	}
	// 最后一个是空格，则删除掉最后的空格，每个单词处理完后加了个空格，这里再处理下
	if ans[len(ans)-1] == ' ' {
		ans = ans[:len(ans)-1]
	}
	return ans
}

// 直接使用库函数
func ReverseWords2(s string) string {
	if len(s) <= 0 {
		return ""
	}
	ans := ""
	// 按" " split
	ss := strings.Split(s, " ")
	// 反向拼接单词，如果是""，不拼接
	for i := len(ss) -1; i >= 0 ; i-- {
		if ss[i] == "" {
			continue
		}
		// 每拼一个单词加个" "
		ans += ss[i] + " "
	}
	// 去掉最后一个空格
	return ans[:len(ans)-1]
}
