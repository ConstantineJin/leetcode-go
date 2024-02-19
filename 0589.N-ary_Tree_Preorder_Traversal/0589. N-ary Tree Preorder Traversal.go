package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ans = []int{root.Val}
	for _, child := range root.Children {
		ans = append(ans, preorder(child)...)
	}
	return ans
}
