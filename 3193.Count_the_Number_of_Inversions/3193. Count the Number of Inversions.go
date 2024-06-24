package main

import "slices"

const mod = 1e9 + 7

func numberOfPermutations(n int, requirements [][]int) int {
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, m+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(int, int) int // dfs(i, j) 表示前 i 个数有 j 个逆序对
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		p := &mem[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if r := req[i-1]; r >= 0 {
			if j < r || j-i > r { // 无法满足
				return 0
			}
			return dfs(i-1, r)
		}
		for k := 0; k <= min(i, j); k++ {
			res += dfs(i-1, j-k)
		}
		return res % mod
	}
	return dfs(n-1, req[n-1])
}
