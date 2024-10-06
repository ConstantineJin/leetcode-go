package main

func constructGridLayout(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	nodes := []int{-1, -1, -1, -1, -1} // nodes[i] 表示节点度数为 i
	for i, to := range g {
		nodes[len(to)] = i // 每个度数的节点记录任意一个
	}
	var row []int
	if nodes[1] != -1 { // 存在度数为 1 的点，必然是 1*n 的情况
		row = []int{nodes[1]}
	} else if nodes[4] == -1 { // 没有度数为 4 的点，必然是 2*n 的情况
		x := nodes[2]
		for _, y := range g[x] {
			if len(g[y]) == 2 {
				row = []int{x, y}
				break
			}
		}
	} else {
		x := nodes[2]
		row = []int{x}
		pre := x
		x = g[x][0]
		for len(g[x]) > 2 { // 边上的节点度数为 3，如果为 2 则到了另一个顶点
			row = append(row, x)
			for _, y := range g[x] {
				if y != pre && len(g[y]) < 4 {
					pre = x
					x = y
					break
				}
			}
		}
		row = append(row, x)
	}
	k := len(row)
	ans := make([][]int, n/k)
	ans[0] = row
	vis := make([]bool, n)
	for _, x := range row {
		vis[x] = true
	}
	for i := 1; i < len(ans); i++ {
		ans[i] = make([]int, k)
		for j, x := range ans[i-1] {
			for _, y := range g[x] {
				if !vis[y] { // x 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
					vis[y] = true
					ans[i][j] = y
					break
				}
			}
		}
	}
	return ans
}
