package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
//func allPossibleFBT(n int) (ans []*TreeNode) {
//	if n%2 == 0 {
//		return
//	} else if n == 1 {
//		ans = append(ans, &TreeNode{Val: 0})
//		return
//	}
//	for i := 1; i < n; i += 2 {
//		var left, right = allPossibleFBT(i), allPossibleFBT(n - i - 1)
//		for _, l := range left {
//			for _, r := range right {
//				ans = append(ans, &TreeNode{Val: 0, Left: l, Right: r})
//			}
//		}
//	}
//	return
//}

// 递推
var dp = [11][]*TreeNode{1: {{}}}

func init() {
	for i := 2; i < 11; i++ { // 计算dp[i]
		for j := 1; j < i; j++ { // 枚举左子树叶子数
			for _, left := range dp[j] { // 枚举左子树
				for _, right := range dp[i-j] { // 枚举右子树
					dp[i] = append(dp[i], &TreeNode{0, left, right})
				}
			}
		}
	}
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 > 0 {
		return dp[(n+1)/2]
	}
	return nil
}
