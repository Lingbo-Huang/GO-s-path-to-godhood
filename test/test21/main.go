package main

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-graphviz"
	"log"
)

// StateMachine 表示有限状态机
type StateMachine struct {
	currentState string
}

// Transition 根据事件进行状态转换
func (sm *StateMachine) Transition(event string) {
	switch sm.currentState {
	case "Idle":
		switch event {
		case "Start":
			sm.currentState = "Running"
		default:
			fmt.Println("Invalid transition from Idle for event", event)
		}

	case "Running":
		switch event {
		case "Pause":
			sm.currentState = "Paused"
		case "Stop":
			sm.currentState = "Stopped"
		default:
			fmt.Println("Invalid transition from Running for event", event)
		}

	case "Paused":
		switch event {
		case "Resume":
			sm.currentState = "Running"
		case "Stop":
			sm.currentState = "Stopped"
		default:
			fmt.Println("Invalid transition from Paused for event", event)
		}

	case "Stopped":
		fmt.Println("Invalid transition from Stopped for any event")

	default:
		fmt.Println("Unknown state")
	}
}

// GetCurrentState 获取当前状态
func (sm *StateMachine) GetCurrentState() string {
	return sm.currentState
}

// VisualizeStateMachine 可视化状态机并显示在窗口中
func VisualizeStateMachine(sm *StateMachine) {
	g := graphviz.New()
	graph, _ := g.Graph()
	defer g.Close()

	// 添加节点
	idleNode, _ := graph.CreateNode("Idle")
	runningNode, _ := graph.CreateNode("Running")
	pausedNode, _ := graph.CreateNode("Paused")
	stoppedNode, _ := graph.CreateNode("Stopped")

	idleNode.SetLabel("Idle")
	idleNode.SetLabel("Running")
	idleNode.SetLabel("Paused")
	idleNode.SetLabel("Stopped")

	// 添加边
	irEdge, _ := graph.CreateEdge("ir", idleNode, runningNode)
	rpEdge, _ := graph.CreateEdge("rp", runningNode, pausedNode)
	rsEdge, _ := graph.CreateEdge("rs", runningNode, stoppedNode)
	prEdge, _ := graph.CreateEdge("pr", pausedNode, runningNode)
	psEdge, _ := graph.CreateEdge("ps", pausedNode, stoppedNode)

	irEdge.SetLabel("Start")
	rpEdge.SetLabel("Pause")
	rsEdge.SetLabel("Stop")
	prEdge.SetLabel("Resume")
	psEdge.SetLabel("Stop")

	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "./graph.png"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 创建状态机
	fsm := &StateMachine{currentState: "Idle"}
	VisualizeStateMachine(fsm)
}
