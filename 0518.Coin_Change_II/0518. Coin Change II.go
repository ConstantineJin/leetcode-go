package main

import "sort"

// 动态规划，递归
//func change(amount int, coins []int) int {
//	if amount == 0 {
//		return 1
//	}
//	var n = len(coins)
//	sort.Ints(coins)
//	var mem = make([][]int, amount)
//	for i := range mem {
//		mem[i] = make([]int, n)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(s, i int) (res int)
//	dfs = func(s, i int) (res int) {
//		if i == n || s+coins[i] > amount {
//			return 0
//		}
//		if mem[s][i] > -1 {
//			return mem[s][i]
//		}
//		defer func() { mem[s][i] = res }()
//		if s+coins[i] == amount {
//			return 1
//		}
//		return dfs(s, i+1) + dfs(s+coins[i], i)
//	}
//	return dfs(0, 0)
//}

// 动态规划，递推
func change(amount int, coins []int) int {
	var n = len(coins)
	sort.Ints(coins)
	var f = make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	f[0][0] = 1
	for i, coin := range coins {
		for j := 0; j <= amount; j++ {
			if j < coin {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = f[i][j] + f[i+1][j-coin]
			}
		}
	}
	return f[n][amount]
}
