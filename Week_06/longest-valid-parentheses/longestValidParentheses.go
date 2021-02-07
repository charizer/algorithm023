package longest_valid_parentheses

// dp解法
func LongestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	n, max := len(s), 0
	dp := make([]int, n)
	// 第0个字符，无配对
	dp[0] = 0
	// 第1个字符，为2或者0
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
		max = 2
	}else{
		dp[1] = 0
	}
	// dp方程
	// dp[i]= dp[i−2]+2;  s[i] = ')', s[i-1] = '(' 的情况
	// dp[i] = dp[i - 1] + 2 + dp[i - dp[i - 1] - 2]; s[i]=')', s[i-1] = ')'的情况
	for i := 2; i < n; i++ {
		// 为左括号，则不会配对
		if s[i] == '(' {
			dp[i] = 0
		}else{
			// 前一个为左括号，则直接配对
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			}else if dp[i-1] > 0 {
				// 找到配对的地方进行判断
				if (i - dp[i - 1] - 1) >= 0 && s[i - dp[i - 1] - 1] == '(' {
					dp[i] = dp[i - 1] + 2
					if (i - dp[i - 1] - 2) >= 0 {
						dp[i] = dp[i] + dp[i - dp[i - 1] - 2]
					}
				}
			}else{
				dp[i] = 0
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 栈解法
func LongestValidParentheses2(s string) int {
	if len(s) <= 1 {
		return 0
	}
	ans, stack := 0, []int{-1}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		}else{
			stack = stack[0:len(stack)-1]
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