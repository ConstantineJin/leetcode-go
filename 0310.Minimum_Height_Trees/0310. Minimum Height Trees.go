package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	var g, deg = make([][]int, n), make([]int, n)
	for _, edge := range edges {
		var x, y = edge[0], edge[1]
		g[x], g[y] = append(g[x], y), append(g[y], x)
		deg[x]++
		deg[y]++
	}
	var nxt []int
	for i, d := range deg {
		if d == 1 {
			nxt = append(nxt, i)
		}
	}
	var remainNodes = n
	for remainNodes > 2 {
		remainNodes -= len(nxt)
		var cur = nxt
		nxt = nil
		for _, x := range cur {
			for _, y := range g[x] {
				deg[y]--
				if deg[y] == 1 {
					nxt = append(nxt, y)
				}
			}
		}
	}
	return nxt
}
