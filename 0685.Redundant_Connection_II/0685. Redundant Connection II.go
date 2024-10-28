package main

type UFS struct{ parent []int }

func newUFS(n int) *UFS {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UFS{parent: parent}
}

func (u *UFS) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UFS) union(x, y int) {
	u.parent[u.find(x)] = u.find(y)
	return
}

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	ufs := newUFS(n + 1)
	fa := make([]int, n+1)
	for i := range n + 1 {
		fa[i] = i
	}
	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if fa[v] != v { // 当前节点有两个父节点，记录冲突
			conflictEdge = edge
		} else {
			fa[v] = u
			if ufs.find(u) == ufs.find(v) { // 当前两个节点已在同一个连通分量中，记录成环
				cycleEdge = edge
			} else {
				ufs.union(u, v)
			}
		}
	}
	if conflictEdge == nil { // 如果不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
		return cycleEdge
	}
	if cycleEdge != nil { // conflictEdge[1] 有两个父节点。如果在访问到 conflictEdge 之前就已经成环，则附加的边在环路上
		return []int{fa[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge // 否则附加边就是 conflictEdge 本身
}
