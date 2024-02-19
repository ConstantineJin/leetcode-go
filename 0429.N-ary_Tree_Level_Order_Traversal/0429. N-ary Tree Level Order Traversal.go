package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return nil
	}
	var cur = []*Node{root}
	for len(cur) > 0 {
		var arr []int
		var nxt []*Node
		for _, node := range cur {
			arr = append(arr, node.Val)
			nxt = append(nxt, node.Children...)
		}
		ans = append(ans, arr)
		cur = nxt
	}
	return ans
}
