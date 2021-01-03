package merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	// 从尾部插入
	// 用于记录数据放入的位置，
	tail, p1, p2 := m + n - 1, m - 1, n - 1
	// nums1和nums2都没有处理完，则一直处理
	for p1 >= 0 && p2 >= 0 {
		// 插入较大值
		if nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		}else{
			nums1[tail] = nums2[p2]
			p2--
		}
		tail--
	}
	// nums1的元素先处理完，需要继续把p2剩余的元素插入nums1中
	if p1 < 0 {
		for i := p2; i >=0; i-- {
			nums1[i] = nums2[i]
		}
	}
	// nums2的元素先处理完，nums1中的其他元素不用再处理
	// if p2 < 0
}
