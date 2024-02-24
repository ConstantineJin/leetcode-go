package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nums = append(nums, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	sort.Ints(nums)
	var ans = make([][]int, len(queries))
	for i, query := range queries {
		var j = sort.SearchInts(nums, query)
		if j < len(nums) && nums[j] == query {
			ans[i] = []int{query, query}
			continue
		}
		switch j {
		case 0:
			ans[i] = []int{-1, nums[0]}
		case len(nums):
			ans[i] = []int{nums[len(nums)-1], -1}
		default:
			ans[i] = []int{nums[j-1], nums[j]}
		}
	}
	return ans
}
