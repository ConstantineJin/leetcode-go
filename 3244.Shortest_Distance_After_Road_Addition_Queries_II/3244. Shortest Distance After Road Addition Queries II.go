package main

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	fa := make([]int, n-1) // 并查集，考虑两个相邻节点中间的边而不是节点
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	ans := make([]int, 0, len(queries))
	cnt := n - 1 // 并查集连通块个数
	for _, query := range queries {
		l, r := query[0], query[1]-1 // 考虑两个相邻节点中间的边而不是节点
		fr := find(r)
		for i := find(l); i < fr; i = find(i + 1) { // 减少合并操作次数
			fa[i] = fr
			cnt--
		}
		ans = append(ans, cnt) // 连通块的个数就是答案
	}
	return ans
}
