package many_many

import (
	"learngo/designPatterns/producerConsumer/out"
	"time"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)

var done = make(chan struct{})

const TaskNum int64 = 10000

func producer(wo chan<- Task, done chan struct{}) {
	var i int64
	for {
		if i >= TaskNum {
			i = 0
		}
		i++
		t := Task{
			ID: i,
		}
		select {
		case wo <- t:
		case <-done:
			out.Println("生产者退出")
			return
		}
	}
}

func consumer(ro <-chan Task, done chan struct{}) {
	for {
		select {
		case t := <-ro:
			if t.ID != 0 {
				t.run()
			}
		case <-done:
			for t := range ro {
				if t.ID != 0 {
					t.run()
				}
			}
			out.Println("消费者退出")
			return
		}
	}
}

func Exec() {
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)

	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)

	time.Sleep(5 * time.Second)
	close(done)
	close(taskCh)
	out.Println("执行成功")
}
