package main

const mod = 1e9 + 7

// 方法1: 记忆化搜索
//func numberOfStableArrays(zero, one, limit int) int {
//	mem := make([][][2]int, zero+1)
//	for i := range mem {
//		mem[i] = make([][2]int, one+1)
//		for j := range mem[i] {
//			mem[i][j] = [2]int{-1, -1}
//		}
//	}
//	var dfs func(i, j, k int) (res int) // dfs(i,j,k) 表示共有 i 个 0 和 j 个 1，第 i+j 位 (arr[i+j-1]) 填 k 的合法方案数
//	dfs = func(i, j, k int) (res int) {
//		if i == 0 { // 递归边界
//			if k == 1 && j <= limit {
//				return 1
//			}
//			return
//		}
//		if j == 0 { // 递归边界
//			if k == 0 && i <= limit {
//				return 1
//			}
//			return
//		}
//		if v := mem[i][j][k]; v > -1 {
//			return v
//		}
//		if k == 0 {
//			res = (dfs(i-1, j, 0) + dfs(i-1, j, 1)) % mod
//			if i > limit {
//				res = (res - dfs(i-limit-1, j, 1) + mod) % mod // dfs(i-1, j, 0) 中包含了最后连续恰好 limit 位都为 0 的方案，需要减去
//			}
//		} else {
//			res = (dfs(i, j-1, 0) + dfs(i, j-1, 1)) % mod
//			if j > limit {
//				res = (res - dfs(i, j-limit-1, 0) + mod) % mod // dfs(i, j-1, 1) 中包含了最后连续恰好 limit 位都为 1 的方案，需要减去
//			}
//		}
//		mem[i][j][k] = res
//		return
//	}
//	return (dfs(zero, one, 0) + dfs(zero, one, 1)) % mod
//}

// 方法2: 递推
func numberOfStableArrays(zero, one, limit int) (ans int) {
	f := make([][][2]int, zero+1)
	for i := range f {
		f[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(limit, zero); i++ {
		f[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		f[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
			if i > limit {
				f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
			}
			f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
			if j > limit {
				f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
			}
		}
	}
	return (f[zero][one][0] + f[zero][one][1]) % mod
}
