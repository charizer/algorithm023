package longest_valid_parentheses

// dp
func LongestValidParentheses(s string) int {
	ans, n := 0, len(s)
	dp := make([]int, n)
	for i := 1; i < len(s); i++ {
		// 以左括号结尾的子串，则肯定不是有效括号
		if s[i] == '(' {
			dp[i] = 0
		}else{
			// 以右括号结尾，并且前一个字符是左括号, 则dp = dp[i-2]+2
			if s[i-1] == '(' {
				// 当前和前一个字符构成有效括号
				dp[i] = 2
				// 还需再加上dp[i-2]
				if i >= 2 {
					dp[i] += dp[i-2]
				}
			} else {
				// 前一个字符也是右括号,
				// 并且前一个字符之前的字符是有效括号
				if dp[i-1] > 0 {
					// 前一个有效括号组的前一个字符是'(' 与当前的')'匹配
					if i - dp[i-1] - 1 >= 0 && s[i - dp[i-1] - 1] == '(' {
						dp[i] = dp[i-1] + 2
						// 前边可能还有有效括号
						if i - dp[i-1] - 2 >= 0 {
							dp[i] += dp[i - dp[i-1] - 2]
						}
					}
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

//栈
func LongestValidParentheses2(s string) int {
	ans := 0
	// 栈底为不匹配的右括号位置，初始化为-1，很重要，能统一逻辑，不用边界特殊判断
	stack := []int{-1}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		}else{
			stack = stack[0:len(stack)-1]
			// 栈空了，压入当前右括号位置作为新的栈底
			if len(stack) == 0 {
				stack = append(stack, i)
			}
			top := stack[len(stack)-1]
			if ans < i - top {
				ans = i - top
			}
		}
	}
	return ans
}

//错误的方法
func LongestValidParentheses3(s string) int {
	ans := 0
	stack := []int{}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		}else{
			// 这样只能计入最近匹配的有效括号组， s = ")()())" 这种情况下只会得到：2
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[0:len(stack)-1]
				if ans < i - top + 1 {
					ans = i - top + 1
				}
			}
		}
	}
	return ans
}
