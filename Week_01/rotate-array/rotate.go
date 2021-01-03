package rotate_array

//暴力
func Rotate(nums []int, k int)  {
	// k可能是nums的倍数
	k = k % len(nums)
	n := len(nums)
	// 每次右移一位，右移k次
	for i := k; i > 0; i-- {
		// 记录下最后一位数值
		tmp := nums[n-1]
		// 其余元素依次后移
		for j := n -2 ; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		// 第一位放置最后一位元素
		nums[0] = tmp
	}
}

// 一次移动k个元素，先记录下这k个元素，把其他的n-k个依次后移k个位置
// 再将记录的k个元素放在头部，申请O(k)的新空间
func Rotate2(nums []int, k int)  {
	k = k % len(nums)
	n := len(nums)
	tmp := make([]int, k)
	// 记录后k个元素，深拷贝，不然后移前n-k个时会被覆盖
	copy(tmp, nums[n-k:])
	// 后移前n-k个元素
	for i := n - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	// 将记录k个元素放在数组头
	for i := 0; i< k ; i++ {
		nums[i] = tmp[i]
	}
}