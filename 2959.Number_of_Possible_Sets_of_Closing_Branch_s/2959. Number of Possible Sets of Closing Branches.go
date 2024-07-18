package main

import "math"

func numberOfSets(n, maxDistance int, roads [][]int) (ans int) {
	g := make([][]int, n) // g[i][j] 表示 i 和 j 的最短距离
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if j != i { // g[i][i] = 0
				g[i][j] = math.MaxInt / 2 // 防止溢出
			}
		}
	}
	for _, road := range roads {
		x, y, wt := road[0], road[1], road[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
next:
	for s := 0; s < 1<<n; s++ { // 枚举子集
		for i, row := range g {
			if s>>i&1 == 0 {
				continue
			}
			copy(f[i], row)
		}
		for k := range f { // Floyd 算法，只考虑在 s 中的节点
			if s>>k&1 == 0 {
				continue
			}
			for i := range f {
				if s>>i&1 == 0 {
					continue
				}
				for j := range f {
					f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				}
			}
		}
		for i, di := range f { // 判断保留的节点中任意两节点的最短路径长度是否都不超过 maxDistance
			if s>>i&1 == 0 {
				continue
			}
			for j, dij := range di {
				if s>>j&1 > 0 && dij > maxDistance {
					continue next
				}
			}
		}
		ans++
	}
	return
}
