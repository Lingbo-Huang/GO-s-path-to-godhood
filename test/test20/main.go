package main

import (
	"fmt"
)

// State 表示状态的类型
type State int

// 定义不同的状态常量
const (
	StateIdle State = iota
	StateRunning
	StatePaused
	StateStopped
)

// Event 表示状态转换的事件类型
type Event int

// 定义不同的事件常量
const (
	EventStart Event = iota
	EventPause
	EventResume
	EventStop
)

// StateMachine 表示有限状态机
type StateMachine struct {
	currentState State
}

// NewStateMachine 创建一个新的状态机
func NewStateMachine() *StateMachine {
	return &StateMachine{currentState: StateIdle}
}

// Transition 根据事件进行状态转换
func (sm *StateMachine) Transition(event Event) {
	switch sm.currentState {
	case StateIdle:
		switch event {
		case EventStart:
			sm.currentState = StateRunning
		default:
			fmt.Println("Invalid transition from StateIdle for event", event)
		}

	case StateRunning:
		switch event {
		case EventPause:
			sm.currentState = StatePaused
		case EventStop:
			sm.currentState = StateStopped
		default:
			fmt.Println("Invalid transition from StateRunning for event", event)
		}

	case StatePaused:
		switch event {
		case EventResume:
			sm.currentState = StateRunning
		case EventStop:
			sm.currentState = StateStopped
		default:
			fmt.Println("Invalid transition from StatePaused for event", event)
		}

	case StateStopped:
		fmt.Println("Invalid transition from StateStopped for any event")

	default:
		fmt.Println("Unknown state")
	}
}

// GetCurrentState 获取当前状态
func (sm *StateMachine) GetCurrentState() State {
	return sm.currentState
}

func main() {
	// 创建状态机
	fsm := NewStateMachine()

	// 输出初始状态
	fmt.Println("Initial State:", fsm.GetCurrentState())

	// 触发事件，改变状态
	fsm.Transition(EventStart)
	fmt.Println("State after EventStart:", fsm.GetCurrentState())

	fsm.Transition(EventPause)
	fmt.Println("State after EventPause:", fsm.GetCurrentState())

	fsm.Transition(EventResume)
	fmt.Println("State after EventResume:", fsm.GetCurrentState())

	fsm.Transition(EventStop)
	fmt.Println("State after EventStop:", fsm.GetCurrentState())
}
