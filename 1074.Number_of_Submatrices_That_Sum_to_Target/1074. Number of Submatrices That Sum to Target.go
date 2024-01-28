package main

func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	var m, n = len(matrix), len(matrix[0])
	var pre = make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		pre[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for p := 1; p <= i; p++ {
				for q := 1; q <= j; q++ {
					if pre[i][j]-pre[p-1][j]-pre[i][q-1]+pre[p-1][q-1] == target {
						ans++
					}
				}
			}
		}
	}
	return
}
