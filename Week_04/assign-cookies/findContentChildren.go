package assign_cookies

import "sort"

func FindContentChildren(g []int, s []int) int {
	// 胃口和饼干先排序
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	// 贪心算法
	for i, j := 0, 0; i < len(g) && j < len(s);{
		// 能够满足胃口，结果+1， 继续判断下一个孩子和胃口
		if s[j] >= g[i] {
			ans++
			i++
			j++
		// 不能满足当前孩子，则看下一块饼干是否能够满足
		}else{
			j++
		}
	}
	return ans
}
