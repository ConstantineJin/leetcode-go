package main

import (
	"container/heap"
	"math"
)

type edge struct{ to, time int }

func maximalPathQuality(values []int, edges [][]int, maxTime int) (ans int) {
	n := len(values)
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{to: y, time: t})
		g[y] = append(g[y], edge{to: x, time: t})
	}
	// Dijkstra 算法
	distance := make([]int, n)
	for i := 1; i < n; i++ {
		distance[i] = math.MaxInt
	}
	h := hp{{0, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		dx, x := p.dis, p.x
		if dx > distance[x] { // x 之前出堆过
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newDistance := dx + e.time
			if newDistance < distance[y] {
				distance[y] = newDistance // 更新 x 的邻居的最短路
				heap.Push(&h, pair{newDistance, y})
			}
		}
	}

	visited := make([]bool, n)
	visited[0] = true
	var dfs func(x, sumTime, sumValue int)
	dfs = func(x, sumTime, sumValue int) {
		if x == 0 {
			ans = max(ans, sumValue)
		}
		for _, e := range g[x] {
			y, t := e.to, e.time
			if sumTime+t+distance[y] > maxTime {
				continue
			}
			if visited[y] {
				dfs(y, sumTime+t, sumValue)
			} else {
				visited[y] = true
				dfs(y, sumTime+t, sumValue+values[y])
				visited[y] = false
			}
		}
	}
	dfs(0, 0, values[0])
	return
}

type pair struct{ dis, x int }
type hp []pair

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].dis < (*h)[j].dis }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)       { a := *h; v = a[len(a)-1]; *h = a[:len(a)-1]; return }
