package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	currentLevel := []*Node{root}
	for len(currentLevel) > 0 {
		var nextLevel []*Node
		for i, node := range currentLevel {
			if i < len(currentLevel)-1 {
				node.Next = currentLevel[i+1]
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
	}
	return root
	//for len(currentLevel) > 0 {
	//	nextLevel := []*Node{}
	//	for _, node := range currentLevel {
	//		if node.Left!= nil {
	//			node.Left.Next = node.Right
	//			if node.Next!= nil {
	//				node.Right.Next = node.Next.Left
	//			}
	//		}
	//		if node.Left!= nil {
	//			nextLevel = append(nextLevel, node.Left, node.Right)
	//		}
	//	}
	//}
	//return root
}
