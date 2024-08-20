package main

import "math"

func stoneGameII(piles []int) int {
	n := len(piles)
	for i := n - 2; i >= 0; i-- {
		piles[i] += piles[i+1] // 后缀和
	}
	mem := make([][]int, n-1)
	for i := range mem {
		mem[i] = make([]int, (n+1)/4+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, m int) (res int) // dfs(i,j) 表示从第 i 堆石头开始，当前回合玩家所能拿到的最多石头，m 含义与题目描述相同
	dfs = func(i, m int) (res int) {
		if i+m*2 >= n { // 全拿
			return piles[i]
		}
		if v := mem[i][m]; v != -1 {
			return v
		}
		mn := math.MaxInt
		for x := 1; x <= m*2; x++ {
			mn = min(mn, dfs(i+x, max(m, x))) // 对手所能拿的石头越少，自己所能拿的石头越多
		}
		mem[i][m] = piles[i] - mn
		return mem[i][m]
	}
	return dfs(0, 1)
}
