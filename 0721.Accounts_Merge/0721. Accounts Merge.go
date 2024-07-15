package main

import "slices"

func accountsMerge(accounts [][]string) (ans [][]string) {
	n := len(accounts)
	emailToIdx := make(map[string][]int) // key 是邮箱，value 是 accounts[i] 的 i
	for i, account := range accounts {
		for _, email := range account[1:] {
			emailToIdx[email] = append(emailToIdx[email], i)
		}
	}
	vis := make([]bool, n)
	emailSet := make(map[string]struct{})
	var dfs func(i int)
	dfs = func(i int) {
		vis[i] = true
		for _, email := range accounts[i][1:] {
			if _, ok := emailSet[email]; ok {
				continue
			}
			emailSet[email] = struct{}{}
			for _, j := range emailToIdx[email] {
				if !vis[j] {
					dfs(j)
				}
			}
		}
	}
	for i, b := range vis {
		if b {
			continue
		}
		clear(emailSet)
		dfs(i)
		res := make([]string, 1, len(emailSet)+1)
		res[0] = accounts[i][0]
		for email := range emailSet {
			res = append(res, email)
		}
		slices.Sort(res[1:])
		ans = append(ans, res)
	}
	return
}
