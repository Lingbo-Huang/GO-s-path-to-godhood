package one_one

import (
	"learngo/designPatterns/producerConsumer/out"
	"sync"
	"time"
)

// Task任务
type Task struct {
	ID int64
}

// 消费任务
func (t *Task) run() {
	out.Println(t.ID)
}

// Task缓冲池
var taskCh = make(chan Task, 10)

// 生产者需要生产的Task数量
const taskNum int64 = 10000

// 生产者 接收参数：单向写入Task类型的chan
// 创建taskNum个Task，并写入通道。全写完后关闭通道
func producer(wo chan<- Task) {
	var i int64
	for i = 1; i <= taskNum; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}
	close(wo)
}

// 消费者 接收参数：单向读取Task类型的chan
// 遍历通道内容，全部消费（调用消费方法）
func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

// 执行函数 让模型跑起来
// 用sync.WaitGroup来确保生产者任务生产完毕且消费者任务消费完毕
func Exec() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(taskCh)
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		consumer(taskCh)
	}(wg)
	go func() {
		out.NewOut().OutPut()
	}()
	time.Sleep(1 * time.Second)
	wg.Wait()
	out.Println("执行成功")
}
