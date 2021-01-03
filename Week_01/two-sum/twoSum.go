package two_sum

func TwoSum(nums []int, target int) []int {
	// 用于记录出现过的元素及下标
	m := make(map[int]int)
	for idx, num := range nums {
		// 有出现过值为target-num的元素，则直接返回下标
		if v, ok := m[target-num]; ok {
			return []int{v, idx}
		}
		// 记录元素及下标
		m[num] = idx
	}
	return []int{}
}
