package find_minimum_in_rotated_sorted_array

func FindMin(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	left, right := 0, len(nums) -1
	for left < right {
		mid := (left + right) / 2
		// 中间比最右边的还大，则最小值在右边，继续查找右边
		// 需要注意数组已经是升序的特殊情况
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		}else{
			// 否则查找左边
			right = mid
		}
	}
	return nums[left]
}
