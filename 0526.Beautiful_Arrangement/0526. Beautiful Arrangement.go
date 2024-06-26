package main

import "math/bits"

func countArrangement(n int) int {
	mem := make([]int, 1<<n)
	for i := range mem {
		mem[i] = -1
	}
	var dfs func(s int) int
	dfs = func(s int) (res int) {
		if s == 1<<n-1 {
			return 1
		}
		if mem[s] != -1 {
			return mem[s]
		}
		i := bits.OnesCount(uint(s)) + 1
		for j := 1; j <= n; j++ {
			if s>>(j-1)&1 == 0 && (i%j == 0 || j%i == 0) {
				res += dfs(s | 1<<(j-1))
			}
		}
		mem[s] = res
		return
	}
	return dfs(0)
}
