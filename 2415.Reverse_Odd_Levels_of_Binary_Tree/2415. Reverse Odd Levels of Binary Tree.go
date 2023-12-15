package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	var cur, odd = []*TreeNode{root}, false // cur记录当前层节点。root为第0层，故odd初始值为false
	for len(cur) > 0 {
		var nxt []*TreeNode
		for _, node := range cur {
			if node.Left != nil {
				nxt = append(append(nxt, node.Left), node.Right) // 本题为完美二叉树，每一层均满
			}
		}
		if odd { // 如果当前是奇数层
			var n = len(cur)
			for i := 0; i < n/2; i++ {
				cur[i].Val, cur[n-i-1].Val = cur[n-i-1].Val, cur[i].Val // 仅反转值
			}
		}
		cur, odd = nxt, !odd // 遍历下一层，层数奇偶性反转
	}
	return root
}
