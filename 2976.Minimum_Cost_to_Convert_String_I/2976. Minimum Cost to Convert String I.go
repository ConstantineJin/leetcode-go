package main

const inf int = 1e13

func minimumCost(source, target string, original, changed []byte, cost []int) (ans int64) {
	var g [26][26]int
	for i := range g {
		for j := range g[i] {
			if i != j {
				g[i][j] = inf
			}
		}
	}
	for i, cst := range cost {
		x, y := original[i]-'a', changed[i]-'a'
		g[x][y] = min(g[x][y], cst)
	}
	// Floyd 算法
	for k := range g {
		for i := range g {
			for j := range g {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
	for i, src := range source {
		ans += int64(g[src-'a'][target[i]-'a'])
	}
	if ans >= int64(inf) {
		return -1
	}
	return
}
