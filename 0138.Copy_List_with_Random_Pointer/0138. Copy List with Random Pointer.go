package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cache = make(map[*Node]*Node)

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if nd, ok := cache[node]; ok {
		return nd
	}
	var newNode = &Node{Val: node.Val}
	cache[node] = newNode
	newNode.Next, newNode.Random = deepCopy(node.Next), deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	clear(cache)
	return deepCopy(head)
}
