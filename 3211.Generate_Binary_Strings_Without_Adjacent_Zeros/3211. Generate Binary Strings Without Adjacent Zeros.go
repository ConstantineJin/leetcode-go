package main

func validStrings(n int) (ans []string) {
	var s []byte
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(s))
			return
		}
		if i == 0 || i > 0 && s[i-1] == '1' {
			s = append(s, '0')
			dfs(i + 1)
			s = s[:i]
		}
		s = append(s, '1')
		dfs(i + 1)
		s = s[:i]
	}
	dfs(0)
	return
}
