package reverse_pairs

func ReversePairs(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	// 分界点
	mid := (left + right) >> 1
	// 分为两部分求和
	count := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	tmp := make([]int, right - left + 1)
	// 合并并计算重要翻转对
	i, t, c := left, left, 0
	for j := mid + 1; j <= right; j, c = j + 1, c + 1 {
		// 不是重要翻转对，i++
		for i <= mid && nums[i] <= 2 * nums[j] {
			i++
		}
		// 归并
		for t <= mid && nums[t] <= nums[j] {
			tmp[c] = nums[t]
			c++
			t++
		}
		tmp[c] = nums[j]
		// 计算重要翻转对
		count += mid - i + 1
	}
	// 将没归并完的进一步归并
	for t <= mid {
		tmp[c] = nums[t]
		c++
		t++
	}
	for i, num := range tmp {
		nums[left+i] = num
	}
	return count
}
