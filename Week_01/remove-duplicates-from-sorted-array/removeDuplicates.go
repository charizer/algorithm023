package remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	// 双指针处理
	// 左指针
	left := 0
	for right := 1; right < len(nums); right++ {
		// 相等，则不处理，右指针继续右移
		if nums[right] == nums[left] {
			continue
		}else{
			// 不相等, 则将元素放入left+1的位置，并更新left
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}
