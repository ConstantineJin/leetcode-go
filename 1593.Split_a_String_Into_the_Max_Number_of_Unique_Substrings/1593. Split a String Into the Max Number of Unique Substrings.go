package main

func maxUniqueSplit(s string) (ans int) {
	n := len(s)
	set := make(map[string]struct{})
	var dfs func(i, cnt int)
	dfs = func(i, cnt int) {
		if n-i+cnt < ans {
			return
		}
		if i == n {
			ans = max(ans, cnt)
			return
		}
		for j := i + 1; j <= n; j++ {
			subStr := s[i:j]
			if _, ok := set[subStr]; !ok {
				set[subStr] = struct{}{}
				dfs(j, cnt+1)
				delete(set, subStr)
			}
		}
	}
	dfs(0, 0)
	return
}
