package search_in_rotated_sorted_array

func Search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	left, right := 0, len(nums) -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[0] <= nums[mid] {
			// 左半部分有序
			// 目标在左半部分
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			}else{
				// 否则目标在右半部分
				left = mid + 1
			}
		}else{
			// 右半部分有序
			// 目标在右半部分
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			}else{
				// 否则目标在左半部分
				right = mid - 1
			}
		}
	}
	return -1
}
