package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
//func dfs(root *TreeNode, cur int) int {
//	if root == nil {
//		return 0
//	}
//	cur = 10*cur + root.Val
//	if root.Left == nil && root.Right == nil {
//		return cur
//	}
//	return dfs(root.Left, cur) + dfs(root.Right, cur)
//}
//
//func sumNumbers(root *TreeNode) int {
//	return dfs(root, 0)
//}

// 数组
func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) []string // 传递[]string而非[]int是为了方便处理连续的0
	dfs = func(node *TreeNode) []string {
		if node == nil {
			return nil
		} else if node.Left == nil && node.Right == nil {
			return []string{strconv.Itoa(node.Val)}
		}
		arr := append(dfs(node.Left), dfs(node.Right)...)
		for i, v := range arr {
			arr[i] = fmt.Sprintf("%d%s", node.Val, v)
		}
		return arr
	}
	for _, s := range dfs(root) {
		v, _ := strconv.Atoi(s)
		ans += v
	}
	return
}
