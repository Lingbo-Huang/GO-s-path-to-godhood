package lru_cache

import (
	"fmt"
	"time"
)

/*
传统的lru会出现预读失败和缓冲池污染的问题

因为预读机制，提前把页放入缓冲池，但有可能最终并没有被读取，即预读失效
新页（例如被预读的页）加入缓冲池时，只加入到老生代头部;
如果数据真正被读取（预读成功），才会加入到新生代的头部;
如果数据没有被读取，则会比新生代里的“热数据页”更早被淘汰出缓冲池

当扫描大量数据时，可能把缓冲池里所有页都替换出去，导致大量热数据被换出
加一个“老生代停留时间窗口T”;
插入老生代头部的页即便被访问，也不会立刻加入新生代头部；
只有同时满足“被访问”，并且“在老生代停留时间大于T”，才会被放入新生代头部
*/

type GenerationType int

const (
	OldGen GenerationType = iota
	NewGen
)

// LRUCacheUpgrade 缓存
type LRUCacheUpgrade struct {
	size             int
	limit            int
	oldLimit         int
	cache            map[string]*DoubleLinkListUpgrade
	headOld, tailOld *DoubleLinkListUpgrade
	headNew, tailNew *DoubleLinkListUpgrade
	windowTime       time.Duration
}

// DoubleLinkListUpgrade 双链表
type DoubleLinkListUpgrade struct {
	key        string
	value      int
	timestamp  time.Time
	generation GenerationType
	next, pre  *DoubleLinkListUpgrade
}

// Get 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1；
func (lc *LRUCacheUpgrade) Get(key string) int {
	if node, ok := lc.cache[key]; !ok {
		return -1
	} else {
		if lc.isOldGen(node) {
			// 如果在老生代，检查停留时间
			if time.Since(node.timestamp) > lc.windowTime {
				lc.moveNodeToNewGen(node)
			} else {
				lc.moveNodeToHead(node, lc.headOld)
			}
		} else {
			lc.moveNodeToHead(node, lc.headNew)
		}
		return node.value
	}
}

// Put 如果关键字 key 存在于缓存中，则变更其数据值 value；如果不存在，则向缓存中插入该 key，value。
// 如果插入操作导致导致 key 数量超过缓存容量，则应该逐出最久未使用的 key；
func (lc *LRUCacheUpgrade) Put(key string, value int, isPrefetch bool) {
	if node, ok := lc.cache[key]; !ok {
		node = &DoubleLinkListUpgrade{
			key:       key,
			value:     value,
			timestamp: time.Now(),
		}
		lc.cache[key] = node
		if isPrefetch {
			lc.addNodeToHead(node, lc.headOld)
		} else {
			lc.addNodeToHead(node, lc.headNew)
		}
		lc.size++
		lc.evict()
	} else {
		node.value = value
		if lc.isOldGen(node) {
			// 如果在老生代，检查停留时间
			if time.Since(node.timestamp) > lc.windowTime {
				lc.moveNodeToNewGen(node)
			} else {
				lc.moveNodeToHead(node, lc.headOld)
			}
		} else {
			lc.moveNodeToHead(node, lc.headNew)
		}
	}
}

// moveNodeToNewGen 将节点从老生代移动到新生代
func (lc *LRUCacheUpgrade) moveNodeToNewGen(node *DoubleLinkListUpgrade) {
	lc.removeNode(node)
	lc.addNodeToHead(node, lc.headNew)
	node.generation = NewGen
}

// moveNodeToHead 将节点移动到链表头部
func (lc *LRUCacheUpgrade) moveNodeToHead(node *DoubleLinkListUpgrade, head *DoubleLinkListUpgrade) {
	lc.removeNode(node)
	lc.addNodeToHead(node, head)
}

// removeNode 从链表中删除节点
func (lc *LRUCacheUpgrade) removeNode(node *DoubleLinkListUpgrade) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// addNodeToHead 将节点添加到指定链表头部
func (lc *LRUCacheUpgrade) addNodeToHead(node *DoubleLinkListUpgrade, head *DoubleLinkListUpgrade) {
	node.next = head.next
	node.pre = head
	head.next.pre = node
	head.next = node
}

// evict 移除超出容量的节点
func (lc *LRUCacheUpgrade) evict() {
	if lc.size <= lc.limit {
		return
	}

	// 首先移除老生代中的节点
	for lc.size > lc.limit && lc.tailOld.pre != lc.headOld {
		lc.removeNode(lc.tailOld.pre)
		delete(lc.cache, lc.tailOld.pre.key)
		lc.size--
	}

	// 然后移除新生代中的节点
	for lc.size > lc.limit && lc.tailNew.pre != lc.headNew {
		lc.removeNode(lc.tailNew.pre)
		delete(lc.cache, lc.tailNew.pre.key)
		lc.size--
	}
}

// isOldGen 判断节点是否在老生代
func (lc *LRUCacheUpgrade) isOldGen(node *DoubleLinkListUpgrade) bool {
	// 如果 node 的前一个节点是老生代的头结点，则 node 属于老生代
	return node.generation == OldGen
}

// ListLRUCache 打印缓存数据
func (lc *LRUCacheUpgrade) ListLRUCache() {
	fmt.Println("Old Generation:")
	node := lc.headOld.next
	for node != lc.tailOld {
		fmt.Printf("key: %s, value: %d\n", node.key, node.value)
		node = node.next
	}
	fmt.Println("New Generation:")
	node = lc.headNew.next
	for node != lc.tailNew {
		fmt.Printf("key: %s, value: %d\n", node.key, node.value)
		node = node.next
	}
}

// NewLRUCacheUpgrade 初始化缓存
func NewLRUCacheUpgrade(limit int, windowTime time.Duration) *LRUCacheUpgrade {
	lc := &LRUCacheUpgrade{
		limit:      limit,
		oldLimit:   limit / 2,
		cache:      make(map[string]*DoubleLinkListUpgrade),
		windowTime: windowTime,
		headOld:    &DoubleLinkListUpgrade{key: "headOld", value: 0, generation: OldGen},
		tailOld:    &DoubleLinkListUpgrade{key: "tailOld", value: 0, generation: OldGen},
		headNew:    &DoubleLinkListUpgrade{key: "headNew", value: 0, generation: NewGen},
		tailNew:    &DoubleLinkListUpgrade{key: "tailNew", value: 0, generation: NewGen},
	}
	lc.headOld.next = lc.tailOld
	lc.tailOld.pre = lc.headOld
	lc.headNew.next = lc.tailNew
	lc.tailNew.pre = lc.headNew
	return lc
}
