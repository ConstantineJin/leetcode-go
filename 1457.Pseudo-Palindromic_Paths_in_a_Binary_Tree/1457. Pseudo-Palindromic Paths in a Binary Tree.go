package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 数组
//func pseudoPalindromicPaths(root *TreeNode) int {
//	p := make([]int, 10)
//	return dfs(root, p)
//}
//
//func dfs(node *TreeNode, p []int) int {
//	if node == nil {
//		return 0
//	}
//	p[node.Val] ^= 1 // 修改 node.Val 出现次数的奇偶性
//	// 返回前，恢复到递归 node 之前的状态（不做这一步就把 node.Val 算到其它路径中了）
//	defer func() { p[node.Val] ^= 1 }()
//	if node.Left == node.Right { // node 是叶子节点
//		if sum(p) <= 1 {
//			return 1
//		}
//		return 0
//	}
//	return dfs(node.Left, p) + dfs(node.Right, p)
//}
//
//func sum(p []int) (s int) {
//	for _, x := range p {
//		s += x
//	}
//	return
//}

func pseudoPalindromicPaths(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, mask int) int {
	if root == nil {
		return 0
	}
	mask ^= 1 << root.Val        // 修改 root.Val 出现次数的奇偶性
	if root.Left == root.Right { // root 是叶子节点
		if mask&(mask-1) == 0 {
			return 1
		}
		return 0
	}
	return dfs(root.Left, mask) + dfs(root.Right, mask)
}
