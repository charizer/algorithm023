package top_k_frequent_elements

import "container/heap"

// 堆中元素内容， key为统计的元素， count为频次
type item struct {
	key int
	count int
}

type hp []item

func (h hp) Len() int {
	return len(h)
}

// 小顶堆
func (h hp) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j],h[i]
}

func (h *hp) Push(ele interface{}) {
	*h = append(*h, ele.(item))
}

func (h *hp) Pop() interface{} {
	ele := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ele
}

func TopKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	m := make(map[int]int)
	// 统计频次
	for _, num := range nums {
		m[num]++
	}
	h := &hp{}
	heap.Init(h)
	// 将map中的数据记录入堆，堆大小限制为k
	for key, val := range m {
		heap.Push(h, item{key: key, count:val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	// 入堆完成后，堆中的k个数据即为出现频次最高的k个数
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).(item).key
	}
	return ans
}
