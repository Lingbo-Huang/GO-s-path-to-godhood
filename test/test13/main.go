package main

import (
	"fmt"
	"time"
)

// Event 结构表示一个事件
type Event struct {
	Message string
}

// EventListener 结构表示事件监听器
type EventListener struct {
	ID       int
	Incoming chan Event
}

// NewEventListener 创建一个新的事件监听器
func NewEventListener(id int) *EventListener {
	return &EventListener{
		ID:       id,
		Incoming: make(chan Event),
	}
}

// Start 启动事件监听器，等待并处理事件
func (el *EventListener) Start() {
	fmt.Printf("Listener %d started\n", el.ID)
	for {
		event := <-el.Incoming // 阻塞等待接收事件
		fmt.Printf("Listener %d received event: %s\n", el.ID, event.Message)
		// 在这里处理事件
	}
}

// EmitEvent 发送事件到监听器
func EmitEvent(listener *EventListener, message string) {
	event := Event{
		Message: message,
	}
	listener.Incoming <- event // 发送事件到通道
}

func main() {
	// 创建两个事件监听器
	listener1 := NewEventListener(1)
	listener2 := NewEventListener(2)

	// 启动事件监听器协程
	go listener1.Start()
	go listener2.Start()

	// 模拟触发事件
	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Event %d", i)
		// 随机选择一个监听器，并发送事件
		if i%2 == 0 {
			EmitEvent(listener1, message)
		} else {
			EmitEvent(listener2, message)
		}
		time.Sleep(time.Second)
	}

	// 注意：在实际应用中，需要考虑关闭监听器协程的机制，以及更复杂的事件处理逻辑。
}
