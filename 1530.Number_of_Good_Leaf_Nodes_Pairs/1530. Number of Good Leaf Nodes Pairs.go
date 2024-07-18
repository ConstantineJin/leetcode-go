package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) (ans int) {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) (leaves []int) {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{0}
		}
		left, right := dfs(node.Left), dfs(node.Right)
		for i := range left {
			left[i]++
		}
		for i := range right {
			right[i]++
		}
		for _, l := range left {
			for _, r := range right {
				if l+r <= distance {
					ans++
				}
			}
		}
		return append(left, right...)
	}
	dfs(root)
	return
}
