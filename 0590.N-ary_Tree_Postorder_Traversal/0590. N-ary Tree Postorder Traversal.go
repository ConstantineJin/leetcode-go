package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (ans []int) {
	if root == nil {
		return nil
	}
	for _, child := range root.Children {
		ans = append(ans, postorder(child)...)
	}
	ans = append(ans, root.Val)
	return ans
}
