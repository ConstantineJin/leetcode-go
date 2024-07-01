package main

type UFS struct {
	parent, size []int
	setCount     int
}

func newUFS(n int) *UFS {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i], size[i] = i, 1
	}
	return &UFS{parent, size, n}
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
		return false
	}
	if u.size[fx] < u.size[fy] {
		fx, fy = fy, fx
	}
	u.size[fx] += u.size[fy]
	u.parent[fy] = fx
	u.setCount--
	return true
}

func (u *UFS) inSameSet(x, y int) bool {
	return u.find(x) == u.find(y)
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	ans := len(edges) // 反向思考，添加尽可能少的边
	alice, bob := newUFS(n), newUFS(n)
	for _, edge := range edges {
		tp, x, y := edge[0], edge[1]-1, edge[2]-1
		if tp == 3 && (!alice.inSameSet(x, y) || !bob.inSameSet(x, y)) { // 优先添加公共边
			alice.union(x, y)
			bob.union(x, y)
			ans--
		}
	}
	ufs := [2]*UFS{alice, bob}
	for _, edge := range edges {
		if tp := edge[0]; tp < 3 && ufs[tp-1].union(edge[1]-1, edge[2]-1) { // 添加这条独占边
			ans--
		}
	}
	if alice.setCount > 1 || bob.setCount > 1 { // 有人无法遍历全图
		return -1
	}
	return ans
}
