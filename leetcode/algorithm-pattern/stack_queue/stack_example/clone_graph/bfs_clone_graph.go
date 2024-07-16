package clone_graph

func cloneGraph_b(node *Node) *Node {
	if node == nil {
		return node
	}
	visited := map[*Node]*Node{}

	queue := []*Node{node}
	visited[node] = &Node{node.Val, []*Node{}}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, neighbor := range n.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &Node{neighbor.Val, []*Node{}}
				queue = append(queue, neighbor)
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
		}
	}
	return visited[node]
}
