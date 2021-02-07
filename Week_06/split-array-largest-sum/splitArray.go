package split_array_largest_sum

func SplitArray(nums []int, m int) int {
	// 将nums分子数组，各子数组和的值的范围
	max, sum := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	left, right := max, sum
	for left < right {
		mid := (right + left) / 2
		// 保证各子数组的值不大于mid，划分子数组
		count := split(nums, mid)
		// 划分出的子数组数量大于m，则mid值小了，在右区间继续处理
		if count > m {
			left = mid + 1
		}else{
			// mid值大了，在左区间继续处理
			right = mid
		}
	}
	return left
}

// 将数组分成count个子数组，每个子数组和小于value
func split(nums []int, value int) int {
	count := 1
	cur := 0
	for _, num := range nums {
		if cur + num > value {
			count++
			cur = num
		}else{
			cur += num
		}
	}
	return count
}
