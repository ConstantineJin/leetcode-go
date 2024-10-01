package main

func mincostTickets(days, costs []int) int {
	n := len(days)
	mem := make([][366]int, n)
	var dfs func(i, expire int) (res int)
	dfs = func(i, expire int) (res int) {
		if i == n || expire >= 365 {
			return
		}
		if mem[i][expire] != 0 {
			return mem[i][expire]
		}
		defer func() { mem[i][expire] = res }()
		if days[i] <= expire {
			return dfs(i+1, expire)
		}
		return min(dfs(i+1, days[i])+costs[0], dfs(i+1, days[i]+6)+costs[1], dfs(i+1, days[i]+29)+costs[2])
	}
	return dfs(0, 0)
}
