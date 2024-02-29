package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	var depth int
	var cur = []*TreeNode{root}
	for len(cur) > 0 {
		var nxt []*TreeNode
		for i, node := range cur {
			if depth%2 == 0 && (node.Val%2 == 0 || i > 0 && node.Val <= cur[i-1].Val) || depth%2 == 1 && (node.Val%2 == 1 || i > 0 && node.Val >= cur[i-1].Val) {
				return false
			}
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
		depth++
	}
	return true
}
