package main

import "math"

const inf = math.MaxInt / 3

type Graph struct{ w [][]int }

// Floyd算法
//func Constructor(n int, edges [][]int) Graph {
//	w := make([][]int, n)
//	for i := range w {
//		w[i] = make([]int, n)
//		for j := range w[i] {
//			if i != j {
//				w[i][j] = inf
//			}
//		}
//	}
//	for _, edge := range edges {
//		w[edge[0]][edge[1]] = edge[2]
//	}
//	for k := range w {
//		for i := range w {
//			for j := range w {
//				w[i][j] = min(w[i][j], w[i][k]+w[k][j])
//			}
//		}
//	}
//	return Graph{w: w}
//}
//
//func (this *Graph) AddEdge(edge []int) {
//	if this.w[edge[0]][edge[1]] <= edge[2] { // 新添加的边的权超过两点间最短路径长度，不可能更新任何最短路径
//		return
//	}
//	for i := range this.w {
//		for j := range this.w {
//			this.w[i][j] = min(this.w[i][j], this.w[i][edge[0]]+edge[2]+this.w[edge[1]][j]) // 枚举所有点对，尝试i-edge[0]-edge[1]-j是否比原来的i-j更小
//		}
//	}
//}
//
//func (this *Graph) ShortestPath(node1 int, node2 int) (d int) {
//	d = this.w[node1][node2]
//	if d == inf {
//		return -1
//	}
//	return
//}

// Dijkstra算法
func Constructor(n int, edges [][]int) Graph {
	w := make([][]int, n)
	for i := range w {
		w[i] = make([]int, n)
		for j := range w[i] {
			if i != j {
				w[i][j] = inf
			}
		}
	}
	for _, edge := range edges {
		w[edge[0]][edge[1]] = edge[2]
	}
	return Graph{w: w}
}

func (this *Graph) AddEdge(edge []int) {
	this.w[edge[0]][edge[1]] = edge[2]
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	n := len(this.w)
	dis, vis := make([]int, n), make([]bool, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[node1] = 0
	for {
		// 找到当前最短路，去更新它的邻居的最短路，
		// 根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == node2 { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range this.w[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
		}
	}
}
