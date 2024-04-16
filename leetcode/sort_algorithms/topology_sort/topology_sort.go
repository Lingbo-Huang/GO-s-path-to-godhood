package main

import "fmt"

// 有向图
type graph struct {
	vertex int           // 顶点数
	list   map[int][]int // 每个顶点到邻接节点列表的映射
}

// 添加从顶点t指向顶点s的边
func (g *graph) addVertex(t int, s int) {
	g.list[t] = push(g.list[t], s)
}

// NewGraph 创建图
func NewGraph(v int) *graph {
	g := new(graph)
	g.vertex = v
	g.list = map[int][]int{}
	for i := 0; i < v; i++ {
		g.list[i] = make([]int, 0)
	}
	return g
}

// 取出切片(队列)第一个元素，返回取出的元素和剩余的切片
func pop(list []int) (int, []int) {
	if len(list) <= 0 {
		return -1, list
	} else {
		a := list[0]
		b := list[1:]
		return a, b
	}
}

// 推入切片（队列）
func push(list []int, value int) []int {
	result := append(list, value)
	return result
}

/* Kahn算法进行拓扑排序
首先计算每个顶点的入度，并将入度为0的顶点加入到队列中。
然后，不断地从队列中取出顶点，输出它，并将其邻接节点的入度减1。
如果邻接节点的入度减为0，将其加入到队列中。重复这个过程，直到队列为空。
*/

func (g *graph) KahnSort() {
	var inDegree = make(map[int]int)
	var queue []int
	for i := 1; i <= g.vertex; i++ {
		for _, m := range g.list[i] {
			inDegree[m]++
		}
	}
	for i := 1; i <= g.vertex; i++ {
		if inDegree[i] == 0 {
			queue = push(queue, i)
		}
	}
	for len(queue) > 0 {
		var now int
		now, queue = pop(queue)
		fmt.Print("-> ", now)
		for _, k := range g.list[now] {
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = push(queue, k)
			}
		}
	}
}

/*
深度优先算法进行拓扑排序
*/
func (g *graph) DfsSort() {
	inverseList := make(map[int][]int)
	// 初始化逆向邻接表(存储每个节点的前驱节点)
	for i := 1; i <= g.vertex; i++ {
		for _, k := range g.list[i] {
			inverseList[k] = append(inverseList[k], i)
		}
	}
	// 初始化访问数组
	visited := make([]bool, g.vertex+1)
	visited[0] = true

	// 对每个未访问过的顶点进行深度优先搜索
	for i := 1; i <= g.vertex; i++ {
		if visited[i] == false {
			visited[i] = true
			dfs(i, inverseList, visited)
		}
	}
}

func dfs(vertex int, inverseList map[int][]int, visited []bool) {
	// 遍历当前顶点的所有前驱节点
	for _, w := range inverseList[vertex] {
		// 如果前驱节点已经被访问过，则跳过
		if visited[w] == true {
			continue
		} else {
			// 否则，递归访问前驱节点
			visited[w] = true
			dfs(w, inverseList, visited)
		}
	}
	fmt.Print("-> ", vertex)
}

func main() {
	g := NewGraph(8)
	g.addVertex(2, 1)
	g.addVertex(3, 1)
	g.addVertex(7, 1)
	g.addVertex(4, 2)
	g.addVertex(5, 2)
	g.addVertex(8, 7)
	g.KahnSort()
	fmt.Println()
	g.DfsSort()
}
