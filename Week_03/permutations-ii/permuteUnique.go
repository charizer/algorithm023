package permutations_ii

import "sort"

func PermuteUnique(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) <= 0 {
		return ans
	}
	// 先进数组排序，方便排列处理
	sort.Ints(nums)
	// 记录已经选过的元素
	used := make([]bool, len(nums))
	backtrack([]int{}, nums, used, &ans)
	return ans
}

// 回溯解法框架
func backtrack(path, nums []int, used []bool, ans *[][]int){
	// 递归终止条件
	if len(path) == len(nums) {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		// 当前位置的数已经选过或者当前数和前一个数相等，但是前一个数还没选，确保相同的数先选前边的再放后边的，不引起重复
		if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		// 选当前数并标记为已选
		path = append(path, nums[i])
		used[i] = true
		backtrack(path, nums, used, ans)
		// 不选当前数并标记为未选
		path = path[:len(path)-1]
		used[i] = false
	}
}

