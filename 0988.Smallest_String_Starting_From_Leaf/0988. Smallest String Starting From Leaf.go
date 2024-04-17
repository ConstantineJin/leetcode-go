package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	var s []string
	var dfs func(node *TreeNode, str string)
	dfs = func(node *TreeNode, str string) {
		str = string(rune(node.Val+'a')) + str
		if node.Left == nil && node.Right == nil {
			s = append(s, str)
		}
		if node.Left != nil {
			dfs(node.Left, str)
		}
		if node.Right != nil {
			dfs(node.Right, str)
		}
	}
	dfs(root, "")
	var ans = s[0]
	for _, str := range s[1:] {
		ans = min(ans, str)
	}
	return ans
}
