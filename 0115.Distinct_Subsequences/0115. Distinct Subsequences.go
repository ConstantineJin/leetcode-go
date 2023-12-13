package main

func numDistinct(s string, t string) int {
	var m, n = len(s), len(t)
	var mem = make([][]int, m+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) (ans int) {
		if j < 0 {
			return 1
		}
		if i < 0 {
			return 0
		}
		var p = &mem[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = ans }()
		if s[i] == t[j] {
			return dfs(i-1, j) + dfs(i-1, j-1)
		} else {
			return dfs(i-1, j)
		}
	}
	return dfs(m-1, n-1) % (1e9 + 7)
}
