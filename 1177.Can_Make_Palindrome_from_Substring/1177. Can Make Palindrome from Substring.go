package main

func canMakePaliQueries(s string, queries [][]int) []bool {
	m, n := len(queries), len(s)
	sum := make([][26]int, n+1)
	for i, c := range s {
		sum[i+1] = sum[i]
		sum[i+1][c-'a']++
	}
	ans := make([]bool, m)
	for i, query := range queries {
		var cnt int
		for j := 0; j < 26; j++ {
			cnt += (sum[query[1]+1][j] - sum[query[0]][j]) % 2
		}
		ans[i] = cnt/2 <= query[2]
	}
	return ans
}
