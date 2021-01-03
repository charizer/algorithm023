package design_circular_deque

type MyCircularDeque struct {
	// 存储数据的slice
	nums []int
	// 队列大小, 方便队列空、满判断，为传入的k+1，多出的这个位置不放元素
	k 	 int
	// 头指针, 队列头即第一个元素所在位置
	head int
	// 尾指针，队列尾，执行最后一个元素的下一个位置
	tail int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		nums: make([]int, k+1),
		k: k+1,
		head: 0,
		tail: 0,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	// 要插入元素先判满
	if this.IsFull() {
		return false
	}
	// 防止越界
	this.head = (this.head - 1 + this.k) % this.k
	this.nums[this.head] = value
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// 要插入元素先判满
	if this.IsFull() {
		return false
	}
	// 尾部插入元素后，尾指针要加一
	this.nums[this.tail] = value
	// 防止越界
	this.tail = (this.tail + 1) % this.k
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	// 要删除元素先判空
	if this.IsEmpty() {
		return false
	}
	// 防止越界
	this.head = (this.head + 1) % this.k
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	// 要删除元素先判空
	if this.IsEmpty() {
		return false
	}
	// 防止越界
	this.tail = (this.tail - 1 + this.k) % this.k
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	// 获取元素先判空
	if this.IsEmpty() {
		return -1
	}
	return this.nums[this.head]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	// 获取元素先判空
	if this.IsEmpty() {
		return -1
	}
	// 防止越界
	return this.nums[(this.tail - 1 + this.k) % this.k]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}else{
		return false
	}
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if (this.tail + 1) % this.k == this.head {
		return true
	}else{
		return false
	}
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
