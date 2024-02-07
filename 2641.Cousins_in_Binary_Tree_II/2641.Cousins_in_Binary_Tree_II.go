package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	var cur = []*TreeNode{root}
	for len(cur) > 0 {
		var nxt []*TreeNode
		var sum int
		for _, node := range cur {
			if node.Left != nil {
				nxt = append(nxt, node.Left)
				sum += node.Left.Val
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
				sum += node.Right.Val
			}
		}
		for _, node := range cur {
			if node.Left != nil && node.Right != nil {
				var temp = sum - node.Left.Val - node.Right.Val
				node.Left.Val, node.Right.Val = temp, temp
			} else if node.Left != nil {
				node.Left.Val = sum - node.Left.Val
			} else if node.Right != nil {
				node.Right.Val = sum - node.Right.Val
			}
		}
		cur = nxt
	}
	root.Val = 0
	return root
}
