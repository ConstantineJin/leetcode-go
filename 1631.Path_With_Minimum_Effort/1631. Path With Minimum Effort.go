package main

import "sort"

type UFS struct {
	parent, size []int
}

func newUFS(n int) *UFS {
	var parent, size = make([]int, n), make([]int, n)
	for i := range parent {
		parent[i], size[i] = i, 1
	}
	return &UFS{parent: parent, size: size}
}

func (u *UFS) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UFS) isSameSet(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *UFS) union(x, y int) {
	var fx, fy = u.find(x), u.find(y)
	if fx == fy {
		return
	}
	if u.size[fx] < u.size[fy] {
		fx, fy = fy, fx
	}
	u.size[fx] += u.size[fy]
	u.parent[fy] = fx
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type edge struct {
	v, w, diff int
}

func minimumEffortPath(heights [][]int) int {
	var edges []edge
	var m, n = len(heights), len(heights[0])
	for i, row := range heights {
		for j, h := range row {
			var id = i*n + j
			if i > 0 {
				edges = append(edges, edge{id - n, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff }) // 对所有的边升序排序
	var ufs = newUFS(m * n)
	for _, e := range edges {
		ufs.union(e.v, e.w)          // 依次加入边
		if ufs.isSameSet(0, m*n-1) { // 如果加入一条权值为x的边后使得左上角到右下角在一个集合中（即连通），该边权值即为答案
			return e.diff
		}
	}
	return 0
}
