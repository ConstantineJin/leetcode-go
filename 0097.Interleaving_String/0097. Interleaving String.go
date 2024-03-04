package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	var m, n = len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	var mem = make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) (res bool) {
		if i == m {
			return s2[j:] == s3[i+j:]
		}
		if j == n {
			return s1[i:] == s3[i+j:]
		}
		switch mem[i][j] {
		case 1:
			return true
		case -1:
			return false
		}
		defer func() {
			if res {
				mem[i][j] = 1
			} else {
				mem[i][j] = -1
			}
		}()
		if s1[i] == s3[i+j] && s2[j] == s3[i+j] {
			return dfs(i+1, j) || dfs(i, j+1)
		}
		if s1[i] == s3[i+j] {
			return dfs(i+1, j)
		}
		if s2[j] == s3[i+j] {
			return dfs(i, j+1)
		}
		return false
	}
	return dfs(0, 0)
}
