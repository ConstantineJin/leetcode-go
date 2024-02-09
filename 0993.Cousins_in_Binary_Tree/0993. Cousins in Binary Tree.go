package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var cur = []*TreeNode{root}
	for len(cur) > 0 {
		var nxt []*TreeNode
		var findX, findY bool
		for _, node := range cur {
			switch node.Val {
			case x:
				findX = true
			case y:
				findY = true
			}
			if node.Left != nil && node.Right != nil && (node.Left.Val == x && node.Right.Val == y || node.Left.Val == y && node.Right.Val == x) {
				return false
			}
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
			cur = nxt
		}
		if findX && findY {
			return true
		} else if findX || findY {
			return false
		}
	}
	return false
}
