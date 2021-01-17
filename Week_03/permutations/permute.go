package permutations

func Permute(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) <= 0 {
		return ans
	}
	backtrack(0, len(nums) - 1, nums, &ans)
	return ans
}

//回溯解法框架
func backtrack(p, q int, nums []int, ans *[][]int) {
	// 递归终止条件
	if p == q {
		*ans = append(*ans, append([]int{},  nums...))
		return
	}
	for i := p; i <= q; i++ {
		// 交换两个数位置
		nums[p], nums[i] = nums[i], nums[p]
		// 继续处理其他数
		backtrack(p + 1, q, nums, ans)
		// 还原两个数位置
		nums[p], nums[i] = nums[i], nums[p]
	}
}


