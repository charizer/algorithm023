学习笔记

1、通过本周学习了解了位运算，能够用位运算解决N皇后问题。

2、熟悉了布隆过滤器及LRU的实现。

3、熟悉了几种常用排序算法的实现。知道了常用几种排序算法的时间复杂度和空间复杂度。

(1)选择排序：基本思想为每一趟从待排序的数据元素中选择最小（或最大）的一个元素作为首元素，直到所有元素排完为止。时间复杂度为O(n^2)

```
func SelectSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[min], nums[i] = nums[i], nums[min]
		}
	}
	return nums
}
```

(2)插入排序: 基本思想是每一步将一个待排序的记录，插入到前面已经排好序的有序序列中去，直到插完所有元素为止。时间复杂度为O(n^2)

```
func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	return nums
}
```

(3)冒泡排序: 基本思想是对相邻的元素进行两两比较，顺序相反则进行交换，这样，每一趟会将最小或最大的元素“浮”到顶端，最终达到完全有序。时间复杂度为O(n^2)
```
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		flag := true
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return nums
}
```

(4)快速排序：基本思想选择一个基准，通过一趟排序将要排序的数据分割成两部分，其中一部分的所有数据都比基准小，另外一部分的所有数据多比基准大，然后再按此方法对这两部分数据递归进行，最后变成有序序列。时间复杂度为O(nlogn)
```
func QuickSort(nums []int) []int {
	left, right := 0, len(nums)-1
	quickSort(left, right, nums)
	return nums
}
func quickSort(left, right int, nums []int) {
	if left < right {
		pivot := partition(left, right, nums)
		quickSort(left, pivot-1, nums)
		quickSort(pivot+1, right, nums)
	}
}
func partition(left, right int, nums []int) int {
	pivot := nums[left]
	for left < right {
		for left < right && pivot <= nums[right] {
			right--
		}
		nums[left] = nums[right]
		for left < right && pivot >= nums[left] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}
```

(5)归并排序：基本思想是分治，分将数组拆分多个子序列，治将子序列合并成较大有序子序列。时间复杂度为O(nlogn)
```
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left, right := 0, len(nums)-1
	mergeSort(left, right, nums)
	return nums
}

func mergeSort(left, right int, nums []int) {
	if left < right {
		mid := (left + right) >> 1
		mergeSort(left, mid, nums)
		mergeSort(mid+1, right, nums)
		merge(left, mid, right, nums)
	}
}

func merge(left, mid, right int, nums []int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for ; i <= mid && j <= right; k++ {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = nums[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		tmp[k] = nums[j]
	}
	for i, num := range tmp {
		nums[left+i] = num
	}
}
```
(6)堆排序: 基本思想是利用堆设计的一种排序算法，时间复杂度均为O(nlogn)。
```
func HeapSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	for i := len(nums)>>1 - 1; i >= 0; i-- {
		heapify(nums, length, i)
	}
	for i := length - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
	return nums
}

func heapify(nums []int, length, i int) {
	left, right, largest := 2*i+1, 2*i+2, i
	if left < length && nums[left] > nums[largest] {
		largest = left
	}
	if right < length && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, length, largest)
	}
}
```
(7) 希尔排序：基本思想是把数组按下标的一定增量分组，对每组使用直接插入排序算法排序；随着增量逐渐减少，每组包含的关键词越来越多，当增量减至1时，整个数组恰被分成一组，算法便终止。时间复杂度O(n^1.3)
```
func ShellSort(nums []int) []int {
	// 数组长度
	n := len(nums)
	// 每次减半，直到步长为 1
	for step := n / 2; step > 0; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i++ {
			for j := i; j-step >= 0 && nums[j] < nums[j-step]; j = j - step {
				nums[j], nums[j-step] = nums[j-step], nums[j]
			}
		}
	}
	return nums
}
```
