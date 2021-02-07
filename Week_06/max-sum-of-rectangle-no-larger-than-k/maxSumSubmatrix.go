package max_sum_of_rectangle_no_larger_than_k

import "math"

func MaxSumSubmatrix(matrix [][]int, k int) int {
	for len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxSum, rows, cols := math.MinInt32, len(matrix), len(matrix[0])
	for left := 0; left < cols; left++ {
		rowSums := make([]int,rows)
		// 扩展一列求一次值，并更新小于k的最大值
		for right := left; right < cols; right++ {
			// 逐行扩展列
			for row := 0; row < rows; row++ {
				rowSums[row] += matrix[row][right]
			}
			// 更新新值
			maxSum = max(maxSum, maxKSubArray(rowSums, k))
			// 提前结束
			if maxSum == k {
				return maxSum
			}
		}
	}
	return maxSum
}

func maxKSubArray(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}
	// O(n)求小于k的子数组
	cur, maxValue := math.MinInt32, 0
	for _, num := range nums {
		if cur > 0 {
			cur = cur + num
		}else{
			cur = num
		}
		// 提前结束
		if cur == k {
			return k
		}
		if cur > maxValue {
			maxValue = cur
		}
	}
	if maxValue < k {
		return maxValue
	}
	// 求得的值比k大，则用暴力方式重新求值O(n^2)
	cur = math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			// 提前结束
			if sum == k {
				return sum
			}
			// 更新小于k的最大值
			if sum < k && sum > cur {
				cur = sum
			}
		}
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



