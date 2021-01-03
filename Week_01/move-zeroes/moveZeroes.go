package move_zeroes

func MoveZeroes(nums []int)  {
	// 用于记录当前非0元素的插入位置
	left := 0
	for right := 0; right < len(nums); right++ {
		// 遍历到的非0元素插入left标记的位置并更新left都下一个位置
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
