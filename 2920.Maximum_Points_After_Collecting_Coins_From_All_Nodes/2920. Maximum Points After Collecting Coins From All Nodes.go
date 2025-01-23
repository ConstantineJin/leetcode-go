package main

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	mem := make([][14]int, n)
	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, fa int) (res int) {
		p := &mem[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res1 := coins[i]>>j - k
		res2 := coins[i] >> (j + 1)
		for _, ch := range g[i] {
			if ch != fa {
				res1 += dfs(ch, j, i) // 不右移
				if j < 13 {           // j+1 >= 14 相当于 res2 += 0 无需递归
					res2 += dfs(ch, j+1, i) // 右移
				}
			}
		}
		return max(res1, res2)
	}
	return dfs(0, 0, -1)
}
