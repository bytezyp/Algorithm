package Lru

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head,tail *DLinkedNode
}

type DLinkedNode struct {
	key,value int
	pre,next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: key,value: value}
}
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		head: initDLinkedNode(0,0),
		tail: initDLinkedNode(0,0),
		cache: make(map[int]*DLinkedNode),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _,ok := this.cache[key]; !ok{
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}
func (this *LRUCache) moveToHead(node *DLinkedNode)  {
	this.removeNode(node)
	this.addToNode(node)
}
func (this *LRUCache) removeNode(node *DLinkedNode)  {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *LRUCache) addToNode(node *DLinkedNode)  {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node // 这行必须在前边
	this.head.next = node
}

func (this *LRUCache) Put(key, value int)  {
	// 新元素
	if _,ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToNode(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed)
			this.size--
		}
	} else{
		// 老元素
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) removeTail() int {
	pre := this.tail.pre
	this.removeNode(pre)
	return pre.key
}






