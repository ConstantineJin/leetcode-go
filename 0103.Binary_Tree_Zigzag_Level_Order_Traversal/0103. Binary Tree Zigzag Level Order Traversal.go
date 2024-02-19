package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return nil
	}
	var cur = []*TreeNode{root}
	for i := 0; len(cur) > 0; i++ {
		var nxt []*TreeNode
		var arr []int
		for _, node := range cur {
			arr = append(arr, node.Val)
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		if i%2 == 1 {
			slices.Reverse(arr)
		}
		ans = append(ans, arr)
		cur = nxt
	}
	return
}
