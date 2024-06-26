package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var bst2arr func(node *TreeNode) (arr []int)
	bst2arr = func(node *TreeNode) (arr []int) {
		if node == nil {
			return
		}
		return append(append(bst2arr(node.Left), node.Val), bst2arr(node.Right)...)
	}
	arr := bst2arr(root)
	var arr2bst func(left, right int) *TreeNode
	arr2bst = func(left, right int) *TreeNode {
		if left > right {
			return nil
		} else if left == right {
			return &TreeNode{Val: arr[left]}
		} else {
			mid := (left + right) / 2
			return &TreeNode{Val: arr[mid], Left: arr2bst(left, mid-1), Right: arr2bst(mid+1, right)}
		}
	}
	return arr2bst(0, len(arr)-1)
}
