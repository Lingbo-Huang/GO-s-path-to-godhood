package main

import (
	"fmt"
	"sync"
)

/*
普通的map里对并发读写直接报错
if h.flags&hashWriting != 0 {
	fatal("concurrent map read and map write")
}
*/

type StringMap struct {
	m sync.Map
}

func (s *StringMap) Store(key, value string) {
	s.m.Store(key, value)
}

func (s *StringMap) Load(key string) (value string, ok bool) {
	v, ok := s.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

func main() {
	var mapex *StringMap = new(StringMap)
	mapex.Store("name", "CaiXiaobo")
	v, _ := mapex.Load("name")
	fmt.Println(v)
}
