package lru_cache

import "fmt"

/*
将最近使用的元素存放到靠近缓存顶部的位置，当一个新条目被访问时，LRU 将它放置到缓存的顶部。
当缓存满时，较早之前访问的条目将从缓存底部被移除。
hash map缓存数据
双链表会更加方便操作尾结点
*/

// LRUCache 缓存
type LRUCache struct {
	size       int
	limit      int
	cache      map[string]*DoubleLinkList
	head, tail *DoubleLinkList
}

// DoubleLinkList 双链表
type DoubleLinkList struct {
	key       string
	value     int
	next, pre *DoubleLinkList
}

// Get 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1；
func (lc *LRUCache) Get(key string) int {
	if _, ok := lc.cache[key]; !ok {
		return -1
	}
	node := lc.cache[key]
	lc.moveNodeToHead(node)
	return node.value
}

// Put 如果关键字 key 存在于缓存中，则变更其数据值 value；如果不存在，则向缓存中插入该 key，value。
// 如果插入操作导致导致 key 数量超过缓存容量，则应该逐出最久未使用的 key；
func (lc *LRUCache) Put(key string, value int) {
	if node, ok := lc.cache[key]; !ok {
		node = &DoubleLinkList{
			key:   key,
			value: value,
		}
		lc.cache[key] = node
		lc.addNodeToHead(node)
		lc.size++
		if lc.size > lc.limit {
			tailPre := lc.tail.pre
			lc.removeNode(tailPre)
			delete(lc.cache, tailPre.key)
			lc.size--
		}
	} else {
		node := lc.cache[key]
		node.value = value
		lc.moveNodeToHead(node)
	}
}

// moveNodeToHead 将节点移动到头部
func (lc *LRUCache) moveNodeToHead(node *DoubleLinkList) {
	lc.removeNode(node)
	lc.addNodeToHead(node)
}

// removeNode 删除节点
func (lc *LRUCache) removeNode(node *DoubleLinkList) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// addNodeToHead 将节点添加到头部
func (lc *LRUCache) addNodeToHead(node *DoubleLinkList) {
	node.next = lc.head.next
	node.pre = lc.head
	lc.head.next.pre = node
	lc.head.next = node
}

// 由于测试需要，实现一个打印 LRUCache 数据的函数和一个初始化 LRUCache 的函数

// ListLRUCache 打印缓存数据
func (lc *LRUCache) ListLRUCache() {
	node := lc.head.next
	for node != nil {
		fmt.Printf("key: %s, value: %d\n", node.key, node.value)
		node = node.next
	}
}

// NewLRUCache 初始化缓存
func NewLRUCache() *LRUCache {
	lc := &LRUCache{
		limit: 5,
		cache: make(map[string]*DoubleLinkList),
		head: &DoubleLinkList{
			key:   "head",
			value: 0,
		},
		tail: &DoubleLinkList{
			key:   "tail",
			value: 0,
		},
	}
	lc.head.next = lc.tail
	lc.tail.pre = lc.head
	return lc
}
