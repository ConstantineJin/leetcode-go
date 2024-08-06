package main

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	from := make([][]int, n) // 终点为当前节点的边的起点
	f := make([]int, n)      // f[i] 表示从 0 到 i 的最短路
	for i := 1; i < n; i++ {
		f[i] = i
	}
	ans := make([]int, 0, len(queries))
	for _, q := range queries {
		l, r := q[0], q[1]
		from[r] = append(from[r], l)
		if f[l]+1 < f[r] { // 可以更新最短路
			f[r] = f[l] + 1
			for i := r + 1; i < n; i++ {
				f[i] = min(f[i], f[i-1]+1)
				for _, j := range from[i] {
					f[i] = min(f[i], f[j]+1)
				}
			}
		}
		ans = append(ans, f[n-1])
	}
	return ans
}
