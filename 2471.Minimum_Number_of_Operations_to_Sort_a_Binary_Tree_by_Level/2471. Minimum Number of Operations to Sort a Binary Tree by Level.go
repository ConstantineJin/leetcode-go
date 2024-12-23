package main

import "sort"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minimumOperations(root *TreeNode) (ans int) {
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		n := len(cur)
		arr := make([]int, n)
		var nxt []*TreeNode
		for i, node := range cur {
			arr[i] = node.Val
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
		// 寻找置换环。对于每一个置换环，交换次数为环的大小减一。环与环之间不发生交换。
		id := make([]int, n) // 离散化
		for i := range id {
			id[i] = i
		}
		sort.Slice(id, func(i, j int) bool { return arr[id[i]] < arr[id[j]] })
		ans += n
		vis := make([]bool, n)
		for _, v := range id {
			if !vis[v] {
				for ; !vis[v]; v = id[v] {
					vis[v] = true
				}
				ans--
			}
		}
	}
	return
}
