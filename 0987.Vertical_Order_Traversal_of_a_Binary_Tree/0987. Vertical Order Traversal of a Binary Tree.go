package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	y, v int
}

func verticalTraversal(root *TreeNode) [][]int {
	var dfs func(node *TreeNode, x, y int)
	var groups = make(map[int][]pair)
	dfs = func(node *TreeNode, x, y int) {
		if node == nil {
			return
		}
		groups[x] = append(groups[x], pair{y, node.Val})
		dfs(node.Left, x-1, y+1)
		dfs(node.Right, x+1, y+1)
	}
	dfs(root, 0, 0)
	var keys = make([]int, 0, len(groups))
	for i := range groups {
		keys = append(keys, i)
	}
	slices.Sort(keys)
	var ans = make([][]int, len(keys))
	for i, key := range keys {
		var g = groups[key]
		slices.SortFunc(g, func(a, b pair) int {
			if a.y != b.y {
				return a.y - b.y
			} else {
				return a.v - b.v
			}
		})
		ans[i] = make([]int, len(g))
		for j, p := range g {
			ans[i][j] = p.v
		}
	}
	return ans
}
