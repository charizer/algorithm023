package trapping_rain_water

func max(a, b int) int {
	if a > b {
		return a
	}else{
		return b
	}
}

// 左右指针法
func Trap(height []int) int {
	if len(height) <= 0 {
		return 0
	}
	n := len(height)
	// 用于记录可存储的水
	sum := 0
	// 设置左、右指针， 左右水高的最大值
	left, right, leftMax, rightMax := 0, n-1, height[0], height[n-1]
	for left < right {
		// 更新左水高最大值
		leftMax = max(height[left], leftMax)
		// 更新右水高最大值
		rightMax = max(height[right],rightMax)
		// 当前位置能盛的水已左右最大水高的较小值来确定
		if leftMax < rightMax {
			 sum += max(leftMax, height[left]) - height[left]
			 left++
		}else{
			sum += max(rightMax, height[right]) - height[right]
			right--
		}
	}
	return sum
}
