package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root *TreeNode
	mp   map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	var mp = make(map[int]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		mp[node.Val] = true
		if node.Left != nil {
			node.Left.Val = node.Val*2 + 1
			dfs(node.Left)
		}
		if node.Right != nil {
			node.Right.Val = node.Val*2 + 2
			dfs(node.Right)
		}
	}
	root.Val = 0
	dfs(root)
	return FindElements{root: root, mp: mp}
}

func (this *FindElements) Find(target int) bool {
	return this.mp[target]
}
