package main

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) (ans []int) {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] })
	var haveSecret = map[int]bool{0: true, firstPerson: true}
	for i, m := 0, len(meetings); i < m; {
		var g, time = make(map[int][]int), meetings[i][2]
		for ; i < m && time == meetings[i][2]; i++ {
			var u, v = meetings[i][0], meetings[i][1]
			g[u], g[v] = append(g[u], v), append(g[v], u)
		}
		var vis = make(map[int]bool)
		var dfs func(v int)
		dfs = func(v int) {
			vis[v], haveSecret[v] = true, true
			for _, u := range g[v] {
				if !vis[u] {
					dfs(u)
				}
			}
		}
		for v := range g {
			if haveSecret[v] && !vis[v] {
				dfs(v)
			}
		}
	}
	for i := range haveSecret {
		ans = append(ans, i)
	}
	sort.Ints(ans)
	return
}
