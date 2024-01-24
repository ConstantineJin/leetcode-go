package main

func maxLength(arr []string) int {
	for i, s := range arr {
		var rec [26]bool
		var repeat = false
		for _, ch := range s {
			if rec[ch-'a'] {
				repeat = true
				break
			}
			rec[ch-'a'] = true
		}
		if repeat {
			arr[i] = ""
		}
	}
	var n = len(arr)
	var rec [26]bool
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}
		var canSelect = true
		for _, ch := range arr[i] {
			if rec[ch-'a'] { // 已经出现过了
				canSelect = false
				break
			}
		}
		var a = dfs(i + 1)
		if canSelect {
			for _, ch := range arr[i] {
				rec[ch-'a'] = true
			}
			defer func() {
				for _, ch := range arr[i] {
					rec[ch-'a'] = false
				}
			}()
			return max(dfs(i+1)+len(arr[i]), a)
		} else {
			return a
		}
	}
	return dfs(0)
}
