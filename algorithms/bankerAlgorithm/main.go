package main

import (
	"fmt"
)

// 定义进程结构体
type Process struct {
	// 进程标识符
	id int
	// 最大需求量
	maximum []int
	// 已分配资源量
	allocated []int
	// 剩余需求量
	need []int
}

// 检查是否可以分配资源给进程
func isSafeState(processes []*Process, available []int, sequence []int) bool {
	// 复制一份可用资源
	work := make([]int, len(available))
	copy(work, available)

	// 创建一个标记，用于记录哪些进程可以被满足
	finished := make([]bool, len(processes))

	// 循环直到所有进程都被满足或者没有可以满足的进程
	for i := 0; i < len(processes); i++ {
		found := false
		for j := 0; j < len(processes); j++ {
			if !finished[j] && canSatisfy(processes[j].need, work) {
				// 可以满足该进程的需求
				found = true
				finished[j] = true
				sequence[i] = j
				// 释放已分配的资源
				work = releaseResources(work, processes[j].allocated)
				break
			}
		}
		// 如果没有找到可以满足的进程，则系统处于不安全状态
		if !found {
			return false
		}
	}
	// 所有进程都可以被满足，系统处于安全状态
	return true
}

// 检查是否可以满足进程的需求
func canSatisfy(need []int, available []int) bool {
	for i := 0; i < len(need); i++ {
		if need[i] > available[i] {
			return false
		}
	}
	return true
}

// 释放已分配的资源
func releaseResources(available []int, allocated []int) []int {
	for i := 0; i < len(allocated); i++ {
		available[i] += allocated[i]
	}
	return available
}

func main() {
	// 定义资源总量
	available := []int{3, 3, 2}

	// 定义进程的最大需求量、已分配资源量和剩余需求量
	processes := []*Process{
		{0, []int{7, 5, 3}, []int{0, 1, 0}, []int{7, 4, 3}},
		{1, []int{3, 2, 2}, []int{2, 0, 0}, []int{1, 2, 2}},
		{2, []int{9, 0, 2}, []int{3, 0, 2}, []int{6, 0, 0}},
		{3, []int{2, 2, 2}, []int{2, 1, 1}, []int{0, 1, 1}},
		{4, []int{4, 3, 3}, []int{0, 0, 2}, []int{4, 3, 1}},
	}

	// 创建一个用于记录安全序列的数组
	sequence := make([]int, len(processes))

	// 检查系统是否处于安全状态
	if isSafeState(processes, available, sequence) {
		fmt.Println("系统处于安全状态")
		fmt.Println("安全序列为:", sequence)
	} else {
		fmt.Println("系统处于不安全状态")
	}
}
