package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var visited = map[*Node]*Node{}
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		var cloneNode = &Node{node.Val, []*Node{}}
		visited[node] = cloneNode
		for _, neighbor := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(neighbor))
		}
		return cloneNode
	}
	return dfs(node)
}
