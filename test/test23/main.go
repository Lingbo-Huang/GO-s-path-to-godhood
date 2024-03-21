package singleton

import (
	"sync"
)

// Singleton 是一个单例结构体
type Singleton struct {
	// 这里可以放置单例需要的其他字段
}

var instance *Singleton
var once sync.Once

// GetInstance 返回单例实例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		// 这里可以进行其他单例初始化操作
	})
	return instance
}
