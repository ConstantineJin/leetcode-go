package main

import "math"

// Floyd算法，递归
//func findTheCity(n int, edges [][]int, distanceThreshold int) (ans int) {
//	w := make([][]int, n) // 边权图
//	for i := range w {
//		w[i] = make([]int, n)
//		for j := range w[i] {
//			w[i][j] = math.MaxInt / 2
//		}
//	}
//	for _, edge := range edges {
//		w[edge[0]][edge[1]], w[edge[1]][edge[0]] = edge[2], edge[2]
//	}
//	memo := make([][][]int, n) // 记忆化搜索
//	for i := range memo {
//		memo[i] = make([][]int, n)
//		for j := range memo[i] {
//			memo[i][j] = make([]int, n)
//		}
//	}
//	var dfs func(k, i, j int) int // dfs(i,j,k)表示i到j的最短路长度，并且该最短路的所有中间节点编号都<=k
//	dfs = func(k, i, j int) int {
//		if k < 0 {
//			return w[i][j]
//		}
//		p := &memo[i][j][k]
//		if *p == 0 {
//			*p = min(dfs(k-1, i, j), dfs(k-1, i, k)+dfs(k-1, k, j)) // 动态规划，不选k与选k两者的较小值
//		}
//		return *p
//	}
//	minCnt := n
//	for i := 0; i < n; i++ {
//		cnt := 0
//		for j := 0; j < n; j++ {
//			if i != j && dfs(n-1, i, j) <= distanceThreshold {
//				cnt++
//			}
//		}
//		if cnt <= minCnt {
//			minCnt, ans = cnt, i
//		}
//	}
//	return
//}

// Floyd算法，递推
//func findTheCity(n int, edges [][]int, distanceThreshold int) (ans int) {
//	w := make([][]int, n) // 边权图
//	for i := range w {
//		w[i] = make([]int, n)
//		for j := range w[i] {
//			w[i][j] = math.MaxInt / 2
//		}
//	}
//	for _, edge := range edges {
//		w[edge[0]][edge[1]], w[edge[1]][edge[0]] = edge[2], edge[2]
//	}
//	f := make([][][]int, n+1) // f[k][i][j]定义同dfs(k,i,j)，表示i到j的最短路长度，并且该最短路的所有中间节点编号都<=k
//	for i := range f {
//		f[i] = make([][]int, n)
//		for j := range f[i] {
//			f[i][j] = make([]int, n)
//		}
//	}
//	f[0] = w // 递归边界，占用f[0]导致后续k需要后移一位
//	for k := 0; k < n; k++ {
//		for i := 0; i < n; i++ {
//			for j := 0; j < n; j++ {
//				f[k+1][i][j] = min(f[k][i][j], f[k][i][k]+f[k][k][j]) // 递推式同递归的状态转移方程
//			}
//		}
//	}
//	minCnt := n
//	for i, dis := range f[n] {
//		cnt := 0
//		for j, d := range dis {
//			if j != i && d <= distanceThreshold {
//				cnt++
//			}
//		}
//		if cnt <= minCnt { // 相等时取最大的 i
//			minCnt, ans = cnt, i
//		}
//	}
//	return
//}

// Floyd算法，递推空间优化
func findTheCity(n int, edges [][]int, distanceThreshold int) (ans int) {
	w := make([][]int, n) // 边权图
	for i := range w {
		w[i] = make([]int, n)
		for j := range w[i] {
			w[i][j] = math.MaxInt / 2
		}
	}
	for _, edge := range edges {
		w[edge[0]][edge[1]], w[edge[1]][edge[0]] = edge[2], edge[2]
	}
	f := w
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j]) // f[k+1][i][k]=f[k][i][k], f[k+1][k][j]=f[k][k][j]，因此第一维可以舍弃
			}
		}
	}
	minCnt := n
	for i, dis := range f {
		cnt := 0
		for j, d := range dis {
			if j != i && d <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= minCnt {
			minCnt, ans = cnt, i
		}
	}
	return
}
