package main

import "sort"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	var arr []int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left), dfs(node.Right)
		if left == -1 || right == -1 || left != right {
			return -1
		}
		arr = append(arr, left+1)
		return left + 1
	}
	dfs(root)
	if len(arr) < k {
		return -1
	}
	sort.Ints(arr)
	return 1<<arr[len(arr)-k] - 1
}
