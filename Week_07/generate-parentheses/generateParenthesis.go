package generate_parentheses

func GenerateParenthesis(n int) []string {
	ans := []string{}
	if n <= 0 {
		return ans
	}
	dfs(&ans, n, 0, 0, "")
	return ans
}

func dfs(ans *[]string, n, left, right int, s string){
	// 左、右括号都到达n, 则是一个符合条件的括号对
	if left == n && right == n {
		*ans = append(*ans, s)
		return
	}
	// 左括号数量 < n, 放一个左括号
	if left < n {
		dfs(ans, n, left + 1, right, s + "(")
	}
	// 右括号数量 < 左括号数量， 放一个右括号
	if right < left {
		dfs(ans, n, left, right + 1, s + ")")
	}
}
