package relative_sort_array

import "sort"

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	ans := []int{}
	// 获取最大值
	upper := 0
	for _, num := range arr1 {
		if num > upper {
			upper = num
		}
	}
	// 计数排序
	freq := make([]int, upper + 1)
	for _, num := range arr1 {
		freq[num]++
	}
	// arr2中的元素按照arr2的顺序放在头部
	for _, num := range arr2 {
		for ;freq[num] > 0; freq[num]-- {
			ans = append(ans, num)
		}
	}
	// 其他元素按照大小放入结果
	for num, count := range freq {
		for ; count > 0; count-- {
			ans = append(ans, num)
		}
	}
	return ans
}

// 自定义排序
func RelativeSortArray2(arr1 []int, arr2 []int) []int {
	// 记录arr2中数值的序号
	rankMap := make(map[int]int)
	for i, num := range arr2 {
		rankMap[num] = i
	}
	// 自定义排序，如果arr2中存在，按照arr2中顺序排序，如果不存在，则按数值排序
	sort.Slice(arr1, func(i, j int) bool {
		rankI, hasI := rankMap[arr1[i]]
		rankJ, hasJ := rankMap[arr1[j]]
		if hasI && hasJ {
			return rankI < rankJ
		}else if hasI || hasJ {
			return hasI
		}else {
			return arr1[i] < arr1[j]
		}
	})
	return arr1
}
