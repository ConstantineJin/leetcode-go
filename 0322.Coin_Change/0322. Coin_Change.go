package main

import "math"

// 动态规划，完全背包，递归
//func coinChange(coins []int, amount int) int {
//	var n = len(coins) // 硬币面值数量
//	var mem = make([][]int, n)
//	for i := range mem {
//		mem[i] = make([]int, amount+1)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(int, int) int
//	dfs = func(i, c int) (res int) {
//		if i < 0 { // 如果没有硬币可供选择
//			if c == 0 { // 如果此时恰好剩余金额也为0
//				return 0 // 方案数为0
//			}
//			return math.MaxInt / 2 // 没有符合的方案
//		}
//		var p = &mem[i][c]
//		if *p != -1 {
//			return *p
//		}
//		defer func() { *p = res }()
//		if c < coins[i] { // 如果当前硬币面值超过剩余金额
//			return dfs(i-1, c) // 不能选择当前的硬币
//		}
//		return min(dfs(i-1, c), dfs(i, c-coins[i])+1) // 选或不选当前面值硬币两种方案的较小值（如果选，后面还可以再选同样面额的）
//	}
//	var ans = dfs(n-1, amount)
//	if ans < math.MaxInt/2 {
//		return ans
//	}
//	return -1
//}

// 递推
//func coinChange(coins []int, amount int) (ans int) {
//	var n = len(coins)
//	var f = make([][]int, n+1)
//	for i := range f {
//		f[i] = make([]int, amount+1)
//	}
//	for i := range f[0] {
//		f[0][i] = math.MaxInt / 2
//	}
//	f[0][0] = 0
//	for i, coin := range coins {
//		for c := 0; c <= amount; c++ {
//			if c < coin {
//				f[i+1][c] = f[i][c]
//			} else {
//				f[i+1][c] = min(f[i][c], f[i+1][c-coin]+1)
//			}
//		}
//	}
//	ans = f[n][amount]
//	if ans < math.MaxInt/2 {
//		return
//	}
//	return -1
//}

// 递推空间优化
func coinChange(coins []int, amount int) (ans int) {
	var n = len(coins)
	var f = [2][]int{make([]int, amount+1), make([]int, amount+1)}
	for i := range f[0] {
		f[0][i] = math.MaxInt / 2
	}
	f[0][0] = 0
	for i, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c < coin {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = min(f[i%2][c], f[(i+1)%2][c-coin]+1)
			}
		}
	}
	ans = f[n%2][amount]
	if ans < math.MaxInt/2 {
		return
	}
	return -1
}
