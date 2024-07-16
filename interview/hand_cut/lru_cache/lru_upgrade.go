package lru_cache

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

type UpgradeLRUCache struct {
	size                               int
	limit                              int
	newSize                            int
	oldSize                            int
	newLimit                           int
	oldLimit                           int
	timeWin                            int
	cache                              map[string]*DoubleList
	newHead, newTail, oldHead, oldTail *DoubleList
}

type DoubleList struct {
	key       string
	value     int
	t         int // 在老生代停留时间
	pre, next *DoubleList
}

// PreRead 预读
func (ulc *UpgradeLRUCache) PreRead(nodes ...*DoubleList) {
	for _, node := range nodes {
		ulc.PutNodeToOldHead(node)
	}
}

func (ulc *UpgradeLRUCache) Get(key string) int {
	if _, ok := ulc.cache[key]; !ok {
		return -1
	}
	node := ulc.cache[key]
	ulc.moveNodeToHead(node)
	return node.value
}

func (ulc *UpgradeLRUCache) PutNodeToOldHead(node *DoubleList) {
	if node, ok := ulc.cache[key]; !ok {

	}
}
