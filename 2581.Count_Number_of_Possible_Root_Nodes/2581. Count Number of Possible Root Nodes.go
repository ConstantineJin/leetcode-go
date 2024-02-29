package main

func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	var n = len(edges) // 边数
	var g = make([][]int, n+1)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	type pair struct{ x, y int }
	var s = make(map[pair]int, n)
	for _, guess := range guesses {
		s[pair{guess[0], guess[1]}] = 1
	}
	var cnt0 int
	var dfs func(x, fa int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以0为根时猜对了一条边
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	var reroot func(x, fa, cnt int) // 换根
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 以x为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}]) // 节点x与y相邻，根从x变成y时只有(x,y)的父子关系倒置，其他无变化
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}
