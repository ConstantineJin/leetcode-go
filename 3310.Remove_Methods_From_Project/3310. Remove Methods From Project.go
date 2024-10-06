package main

func remainingMethods(n, k int, invocations [][]int) (ans []int) {
	g := make([][]int, n)
	for _, invocation := range invocations {
		g[invocation[0]] = append(g[invocation[0]], invocation[1])
	}
	vis := make([]bool, n)
	vis[k] = true
	cur := []int{k}
	for len(cur) > 0 {
		var nxt []int
		for _, method := range cur {
			for _, m := range g[method] {
				if !vis[m] {
					nxt = append(nxt, m)
					vis[m] = true
				}
			}
		}
		cur = nxt
	}
	for _, invocation := range invocations {
		if vis[invocation[1]] && !vis[invocation[0]] {
			ans = make([]int, n)
			for i := range ans {
				ans[i] = i
			}
			return ans
		}
	}
	for i := range n {
		if !vis[i] {
			ans = append(ans, i)
		}
	}
	return
}
