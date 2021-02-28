package lru_cache

// 双链表节点定义
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		size: 0,
		capacity: capacity,
		cache: make(map[int]*DLinkedNode),
		head: &DLinkedNode{key: 0,value: 0},
		tail: &DLinkedNode{key: 0, value: 0},
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

// 获取元素
func (this *LRUCache) Get(key int) int {
	// 不存在，返回-1
	if node, ok := this.cache[key]; !ok {
		return -1
	}else{
		// 存在，移到头部
		this.moveToHead(node)
		return node.value
	}
}

// 插入元素
func (this *LRUCache) Put(key int, value int)  {
	// 已经存在
	if node, ok := this.cache[key]; ok {
		node.value = value
		// 将节点移到首部
		this.moveToHead(node)
	}else{
		// 已经满了
		if this.size == this.capacity {
			//把尾部节点删了
			remove := this.removeTail()
			delete(this.cache, remove.key)
			this.size--
		}
		node = &DLinkedNode{key: key, value: value}
		this.addToHead(node)
		this.size++
		this.cache[key] = node
	}
}

// 将节点插入双链表头部
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 删除节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 将节点移到链表头部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除尾部节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */