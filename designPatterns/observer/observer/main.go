package main

import "fmt"

// 主题接口
type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify()
}

// 观察者接口
type Observer interface {
	Update()
}

// 具体主题
type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Remove(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify() {
	for _, observer := range s.observers {
		observer.Update()
	}
}

// 具体观察者
type ConcreteObserver struct{}

func (o *ConcreteObserver) Update() {
	fmt.Println("具体观察者收到更新通知")
}

func main() {
	subject := &ConcreteSubject{}
	observer := &ConcreteObserver{}

	subject.Register(observer)
	subject.Notify()
}
