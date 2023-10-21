package main

// UFSet 并查集
type UFSet struct {
	parents []int
	sizes   []int
}

// NewUFSet 为构造函数
func NewUFSet(n int) *UFSet {
	parents, sizes := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		parents[i], sizes[i] = i, 1
	}
	return &UFSet{parents: parents, sizes: sizes}
}

// Find 获取元素所在并查集的根节点
func (ufs *UFSet) Find(x int) int {
	if ufs.parents[x] == x {
		return x
	}
	ufs.parents[x] = ufs.Find(ufs.parents[x])
	return ufs.parents[x]
}

// Union 合并两个元素所在的并查集
func (ufs *UFSet) Union(x, y int) {
	xRoot, yRoot := ufs.Find(x), ufs.Find(y)
	if xRoot != yRoot {
		if ufs.sizes[x] > ufs.sizes[y] {
			ufs.parents[yRoot], ufs.sizes[xRoot] = xRoot, ufs.sizes[xRoot]+ufs.sizes[yRoot]
		} else {
			ufs.parents[xRoot], ufs.sizes[yRoot] = yRoot, ufs.sizes[xRoot]+ufs.sizes[yRoot]
		}
	}
}

// GetSize 获取元素所在并查集的大小
func (ufs *UFSet) GetSize(x int) int {
	return ufs.sizes[x]
}

func countPairs(n int, edges [][]int) (res int64) {
	ufs := NewUFSet(n)
	for _, edge := range edges {
		ufs.Union(edge[0], edge[1])
	}
	for i := 0; i < n; i++ {
		res += int64(n - ufs.GetSize(ufs.Find(i)))
	}
	return res / 2
}

//// DFS
//func countPairs(n int, edges [][]int) (res int64) {
//	graph := make([][]int, n)
//	for _, edge := range edges {
//		graph[edge[0]] = append(graph[edge[0]], edge[1])
//		graph[edge[1]] = append(graph[edge[1]], edge[0])
//	}
//	vis := make([]bool, n)
//	var dfs func(int) int
//	dfs = func(i int) int {
//		if vis[i] {
//			return 0
//		}
//		vis[i] = true
//		cnt := 1
//		for _, j := range graph[i] {
//			cnt += dfs(j)
//		}
//		return cnt
//	}
//	var s int64
//	for i := 0; i < n; i++ {
//		t := int64(dfs(i))
//		res += s * t
//		s += t
//	}
//	return
//}
