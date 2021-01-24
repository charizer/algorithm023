package jump_game

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CanJump(nums []int) bool {
	curMax := 0
	for i := 0; i < len(nums); i++ {
		// 当前位置比上次能跳到的最大位置还大，则不能跳到位置i
		if  i > curMax {
			return false
		}
		// 更新当前能跳到的最大位置
		curMax = max(curMax, i + nums[i])
	}
	return true
}
