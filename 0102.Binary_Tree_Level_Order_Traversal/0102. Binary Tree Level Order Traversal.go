package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return nil
	}
	var cur = []*TreeNode{root}
	for len(cur) > 0 {
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
		ans = append(ans, arr)
		cur = nxt
	}
	return
}
