package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes, isChild := make(map[int]*TreeNode), make(map[int]bool)
	for _, description := range descriptions {
		parent, child := description[0], description[1]
		if nodes[parent] == nil {
			nodes[parent] = &TreeNode{Val: parent}
		}
		if nodes[child] == nil {
			nodes[child] = &TreeNode{Val: child}
		}
		if description[2] == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
		isChild[child] = true
	}
	for i, node := range nodes {
		if !isChild[i] {
			return node
		}
	}
	return nil
}
