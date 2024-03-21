package many_one

import (
	"learngo/designPatterns/producerConsumer/out"
	"sync"
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

// 每个生产者需要生产的任务数量
const nums int64 = 100

// 生产者 接收参数：单向写入Task类型的chan，当前任务编号startNum, 该生产者需要生产的任务数nums
// 从当前任务编号开始创建nums个Task，并写入通道。全写完后关闭通道
func producer(wo chan<- Task, startNum int64, nums int64) {
	var i int64
	for i = startNum; i < (startNum + nums); i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}
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
	//保证生产者生产完毕
	wg := &sync.WaitGroup{}
	//保证生产者生产完毕后，channel关闭
	pwg := &sync.WaitGroup{}
	var i int64
	wg.Add(1)
	for i = 0; i < taskNum; i += nums {
		if i >= taskNum {
			break
		}
		wg.Add(1)
		pwg.Add(1)
		go func(i int64) {
			defer wg.Done()
			defer pwg.Done()
			producer(taskCh, i, nums)
		}(i)
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		consumer(taskCh)
	}(wg)

	pwg.Wait()
	go close(taskCh)

	wg.Wait()
	out.Println("执行成功")
}
