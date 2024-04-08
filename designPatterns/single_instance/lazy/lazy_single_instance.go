package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

// 锁
var lock sync.Mutex

// 标记
var initialized uint32

var instance *singleton

func GetInstance() *singleton {
	// 如果已经初始化过，则直接返回该单例
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 为了线程安全，防止多个goroutine同时调用该方法，都判断为nil，创建对象，就不符合单例了，所以增加互斥

	//如果没有初始化过，则加锁申请
	// 如果只是加普通锁，而没有原子操作标记判断，这样每次创建单例对象都要申请锁释放锁，增加开销
	lock.Lock()
	defer lock.Unlock()

	// 只有首次GetInstance()方法被调用，才会生成这个单例对象
	if instance == nil {
		instance = new(singleton)
		// 设置标记为1
		atomic.StoreUint32(&initialized, 1)
		return instance
	}
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例的某方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
