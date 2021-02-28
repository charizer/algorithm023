package merge_intervals

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Merge(intervals [][]int) [][]int {
	ans := [][]int{}
	if len(intervals) <= 0 {
		return ans
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if ans[len(ans)-1][1] >= intervals[i][0] {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		}else{
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
