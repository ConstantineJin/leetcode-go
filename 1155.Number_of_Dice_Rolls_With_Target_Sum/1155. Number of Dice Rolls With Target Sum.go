package main

// 递归
func numRollsToTarget(n int, k int, target int) int {
	if target < 0 || target > n*k {
		return 0
	}
	var mem = make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, target+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(m, t int) (res int) {
		if m == 0 {
			if t == 0 {
				return 1
			}
			return 0
		}
		var p = &mem[m][t]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		for i := 1; i <= k && i <= t; i++ {
			res = (res + dfs(m-1, t-i)) % (1e9 + 7)
		}
		return
	}
	return dfs(n, target)
}

//func numRollsToTarget(n int, k int, target int) int {
//	const mod int = 1e9 + 7
//	f := make([]int, target+1)
//	f[0] = 1
//	for i := 1; i <= n; i++ {
//		g := make([]int, target+1)
//		for j := 1; j <= min(target, i*k); j++ {
//			for h := 1; h <= min(j, k); h++ {
//				g[j] = (g[j] + f[j-h]) % mod
//			}
//		}
//		f = g
//	}
//	return f[target]
//}
