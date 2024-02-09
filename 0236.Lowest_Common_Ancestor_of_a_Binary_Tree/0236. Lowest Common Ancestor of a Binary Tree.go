package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root == p || root == q {
		return root
	} else {
		left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
		if left != nil && right != nil {
			return root
		} else if left != nil {
			return left
		} else if right != nil {
			return right
		} else {
			return nil
		}
	}
}
