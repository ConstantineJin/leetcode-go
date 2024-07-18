package main

import "container/heap"

type Edge struct{ to, time int }

// Dijkstra 算法
func minimumTime(n int, edges [][]int, disappear []int) []int {
	g := make([][]Edge, n) // 邻接表
	for _, edge := range edges {
		u, v, t := edge[0], edge[1], edge[2]
		g[u] = append(g[u], Edge{v, t})
		g[v] = append(g[v], Edge{u, t})
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		dx, x := p.dis, p.x
		if dx > dis[x] {
			continue
		}
		for _, edge := range g[x] {
			y, newDis := edge.to, edge.time+dx
			if newDis < disappear[y] && (dis[y] < 0 || newDis < dis[y]) {
				dis[y] = newDis
				heap.Push(&h, pair{newDis, y})
			}
		}
	}
	return dis
}

type pair struct{ dis, x int }
type hp []pair

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].dis < (*h)[j].dis }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)       { a := *h; *h, v = (*h)[:len(*h)-1], a[len(a)-1]; return v }
