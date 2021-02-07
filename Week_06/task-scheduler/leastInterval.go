package task_scheduler

import "sort"

/*
	任务分桶
	数量最多的任务先入桶

	n + 1 : 冷却时间 + 1个单位执行时间
	A	B	冷却
	A	B
	A	B
	A	B

	第一种情况：
 		maxTimes：最多任务的次数
		maxCount：与最多任务相同的数量
		sum = (maxTimes−1)∗(n+1)+maxCount

	第二种情况：
		任务多，全部可以填冷却时间
		sum = len(tasks)

	取两个中的较大值
 */


func LeastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	// 任务桶
	buckets := make([]int, 26)
	for _, task := range tasks {
		buckets[task-'A']++
	}
	sort.Ints(buckets)
	// 最大次数的任务及最大次数任务的个数
	maxTimes, maxCount := buckets[25], 1
	for i := 25; i > 0; i-- {
		if buckets[i-1] == buckets[i] {
			maxCount++
		}else{
			break
		}
	}
	return max(len(tasks), (maxTimes - 1)*(n + 1) + maxCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
