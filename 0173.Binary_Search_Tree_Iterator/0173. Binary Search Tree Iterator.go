package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nums []int
	i    int
}

func Constructor(root *TreeNode) BSTIterator {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return nil
		}
		// go version < 1.22
		//var nums = append(dfs(node.Left), node.Val)
		//nums = append(nums, dfs(node.Right)...)
		//return nums
		// go version >= 1.22
		return slices.Concat(dfs(node.Left), []int{node.Val}, dfs(node.Right))
	}
	return BSTIterator{dfs(root), 0}
}

func (this *BSTIterator) Next() int {
	defer func() { this.i++ }()
	return this.nums[this.i]
}

func (this *BSTIterator) HasNext() bool {
	return this.i < len(this.nums)-1
}
