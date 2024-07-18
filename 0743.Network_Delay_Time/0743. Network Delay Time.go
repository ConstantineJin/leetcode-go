package main

import (
	"math"
	"slices"
)

const inf = math.MaxInt / 2

// Dijkstra 算法
func networkDelayTime(times [][]int, n, k int) int {
	g := make([][]int, n) // g[i][j] 表示节点 i 到节点 j 的边权，如果不存在则为 ∞
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, time := range times {
		g[time[0]-1][time[1]-1] = time[2] // 有向图
	}
	dis := make([]int, n) // dis[i] 表示起点 k 到节点 i 的最短路径长度
	for i := range dis {
		dis[i] = inf
	}
	dis[k-1] = 0
	done := make([]bool, n) // 记录是否已求得从节点 k 到节点 i 的最短路径
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 { // 没有节点被更新
			return slices.Max(dis)
		}
		if dis[x] == inf { // 有节点无法到达
			return -1
		}
		done[x] = true
		for y, d := range g[x] { // 更新 x 的邻居的最短路
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
}
