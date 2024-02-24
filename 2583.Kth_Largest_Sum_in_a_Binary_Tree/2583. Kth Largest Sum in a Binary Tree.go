package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var cur, sums = []*TreeNode{root}, make([]int, 0, k)
	for len(cur) > 0 {
		var nxt []*TreeNode
		var sum int
		for _, node := range cur {
			sum += node.Val
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		sums = append(sums, sum)
		cur = nxt
	}
	if len(sums) < k {
		return -1
	}
	sort.Ints(sums)
	return int64(sums[len(sums)-k])
}
