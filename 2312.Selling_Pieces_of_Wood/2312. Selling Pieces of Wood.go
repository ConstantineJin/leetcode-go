package main

func sellingWood(m int, n int, prices [][]int) int64 {
	type pair struct{ x, y int }
	var mp = make(map[pair]int)
	for _, price := range prices {
		mp[pair{price[0], price[1]}] = price[2]
	}
	var f = make([][]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			f[i][j] = mp[pair{i, j}]
			for k := 1; k <= j/2; k++ {
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k <= i/2; k++ {
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return int64(f[m][n])
}
