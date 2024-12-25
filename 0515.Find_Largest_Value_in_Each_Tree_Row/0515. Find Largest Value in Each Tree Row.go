package main

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		var nxt []*TreeNode
		ans = append(ans, math.MinInt)
		for _, node := range cur {
			ans[len(ans)-1] = max(ans[len(ans)-1], node.Val)
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
	}
	return
}
