package main

import "container/heap"

type cell struct{ x, y, h int }

var dirs = [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func trapRainWater(heightMap [][]int) (ans int) {
	m, n := len(heightMap), len(heightMap[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	h := new(hp) // 木桶原理，使用小顶堆维护当前最外围的木板长度
	for i := 0; i < m; i++ {
		heap.Push(h, cell{x: i, y: 0, h: heightMap[i][0]})
		heap.Push(h, cell{x: i, y: n - 1, h: heightMap[i][n-1]})
		vis[i][0], vis[i][n-1] = true, true
	}
	for j := 1; j < n-1; j++ {
		heap.Push(h, cell{x: 0, y: j, h: heightMap[0][j]})
		heap.Push(h, cell{x: m - 1, y: j, h: heightMap[m-1][j]})
		vis[0][j], vis[m-1][j] = true, true
	}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		for _, dir := range dirs {
			x, y := cur.x+dir[0], cur.y+dir[1]
			if 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] {
				ans += max(cur.h-heightMap[x][y], 0)
				vis[x][y] = true
				heap.Push(h, cell{x, y, max(heightMap[x][y], cur.h)})
			}
		}
	}
	return
}

type hp []cell

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].h < (*h)[j].h }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(cell)) }
func (h *hp) Pop() any           { defer func() { *h = (*h)[:len(*h)-1] }(); return (*h)[len(*h)-1] }
