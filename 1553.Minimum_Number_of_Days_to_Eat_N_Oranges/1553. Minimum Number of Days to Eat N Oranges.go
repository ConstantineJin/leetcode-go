package main

func minDays(n int) int {
	mem := make(map[int]int)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i <= 1 {
			return i
		}
		if v, ok := mem[i]; ok {
			return v
		}
		mem[i] = min(dfs(i/2)+i%2, dfs(i/3)+i%3) + 1
		return mem[i]
	}
	return dfs(n)
}
