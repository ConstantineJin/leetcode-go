package main

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- { // 从倒数第二层开始向上遍历
		for j := range triangle[i] { // 每一层中的遍历顺序任意
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1]) // 递推关系：三角形中每个节点到底边的最小路径和等于下一行相同下标和下标+1到底边的最小路径和加上当前节点的值；本题可以原地存储结果
		}
	}
	return triangle[0][0]
}
