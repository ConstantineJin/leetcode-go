package main

import "math"

type Edge struct{ to, id int }

func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
	g := make([][]Edge, n)
	for i, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], Edge{v, i})
		g[v] = append(g[v], Edge{u, i})
	}
	var delta int // target 减去第一遍运行 Dijkstra 算法时从 source 到 destination 的最短路长度
	dis := make([][2]int, n)
	for i := range dis {
		dis[i] = [2]int{math.MaxInt32 >> 1, math.MaxInt32 >> 1}
	}
	dis[source] = [2]int{}
	dijkstra := func(k int) { // k = 0/1 表示第一/二次运行最短路算法
		vis := make([]bool, n)
		for {
			x := -1
			for y, b := range vis {
				if !b && (x < 0 || dis[y][k] < dis[x][k]) {
					x = y
				}
			}
			if x == destination { // 从 source 到 destination 的最短路已确定
				return
			}
			vis[x] = true
			for _, edge := range g[x] {
				y, weight := edge.to, edges[edge.id][2]
				if weight == -1 {
					weight = 1 // 将边权为 -1 的边全部变成 1
				}
				if k == 1 && edges[edge.id][2] == -1 {
					w := delta + dis[y][0] - dis[x][1]
					if w > weight {
						weight = w
						edges[edge.id][2] = w
					}
				}
				dis[y][k] = min(dis[y][k], dis[x][k]+weight) // 更新最短路
			}
		}
	}
	dijkstra(0)
	delta = target - dis[destination][0]
	if delta < 0 { // 如果把全部 -1 改成 1，最短路比 target 还大，则不存在可行方案
		return nil
	}
	dijkstra(1)
	if dis[destination][1] < target { // 最短路无法再变大时，仍无法达到 target，则不存在可行方案
		return nil
	}
	for _, edge := range edges {
		if edge[2] == -1 {
			edge[2] = 1
		}
	}
	return edges
}
