package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache 是一个简单的内存缓存结构
type Cache struct {
	data sync.Map
}

// Get 从缓存中获取数据，如果不存在则计算并存入缓存
func (c *Cache) Get(key string, calculateFunc func() interface{}) interface{} {
	// 从缓存中尝试获取数据
	if value, ok := c.data.Load(key); ok {
		return value
	}

	// 如果缓存中不存在，则计算数据
	result := calculateFunc()

	// 将计算结果存入缓存
	c.data.Store(key, result)

	return result
}

func main() {
	// 创建缓存实例
	cache := &Cache{}

	// 示例：获取数据，如果缓存中不存在则计算
	result1 := cache.Get("key1", func() interface{} {
		fmt.Println("Calculating for key1...")
		time.Sleep(2 * time.Second) // 模拟计算耗时
		return "Result for key1"
	})
	fmt.Println("Result 1:", result1)

	// 再次获取相同的数据，应该从缓存中直接获取而不再计算
	result2 := cache.Get("key1", func() interface{} {
		fmt.Println("Calculating for key1...")
		time.Sleep(2 * time.Second)
		return "Result for key1"
	})
	fmt.Println("Result 2:", result2)
}
