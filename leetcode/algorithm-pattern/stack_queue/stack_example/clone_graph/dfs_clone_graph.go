package clone_graph

/*
给你无向连通图中一个节点的引用，
请你返回该图的深拷贝（克隆）。
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	// visited[node] 非空表示 node 已经被拷贝了, 直接把克隆节点返回
	if v, ok := visited[node]; ok {
		return v
	}

	// 拷贝 Node
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	// 标记 node 已经被拷贝为 newNode
	visited[node] = newNode

	// 拷贝 Node 的Neighbors
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}
