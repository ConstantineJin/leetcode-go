package main

func superEggDrop(k, n int) int {
	f := make([][]int, n+1) // f[i][j] 表示在有 i 次操作机会和 j 枚鸡蛋的情况下，可以确定 f 值的最大建筑层数
	f[0] = make([]int, k+1)
	for i := 1; ; i++ {
		f[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			f[i][j] = f[i-1][j] + f[i-1][j-1] + 1 // 在 f[i-1][j-1]+1 楼扔第一枚鸡蛋，如果碎了转移到 f[i-1][j-1]，如果没碎就转移到 f[i-1][j]
		}
		if f[i][k] >= n {
			return i
		}
	}
}
