package main

type UFS struct {
	parent, size []int
	setCount     int // 当前集合数量，即连通分量数量
}

func newUFS(n int) *UFS {
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i], size[i] = i, 1
	}
	return &UFS{parent: parent, size: size, setCount: n}
}

func (u *UFS) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UFS) union(x, y int) {
	fx, fy := u.find(x), u.find(y)
	if fx == fy {
		return
	}
	if u.size[fx] < u.size[fy] {
		fx, fy = fy, fx
	}
	u.size[fx] += u.size[fy]
	u.parent[fy] = fx
	u.setCount--
}

func findCircleNum(isConnected [][]int) int {
	var ufs = newUFS(len(isConnected))
	for i, row := range isConnected {
		for j, v := range row[i+1:] {
			if v == 1 {
				ufs.union(i, i+j+1)
			}
		}
	}
	return ufs.setCount
}
