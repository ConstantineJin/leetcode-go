package main

type UFS struct{ parent, size []int }

func newUFS(n int) *UFS {
	parent, size := make([]int, n), make([]int, n)
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

func (u *UFS) union(x, y int) bool {
	fx, fy := u.find(x), u.find(y)
	if fx == fy {
		return true
	}
	if u.size[fx] < u.size[fy] {
		fx, fy = fy, fx
	}
	u.size[fx] += u.size[fy]
	u.parent[fy] = fx
	return false
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	ufs := newUFS(n)
	for _, edge := range edges {
		if ufs.union(edge[0]-1, edge[1]-1) {
			return edge
		}
	}
	return nil
}
