package main

func specialPerm(nums []int) (ans int) {
	n := len(nums)
	u := 1<<n - 1
	mem := make([][]int, u)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(s, i int) (res int) {
		if s == 0 {
			return 1 // 找到一个特别排列
		}
		if mem[s][i] != -1 {
			return mem[s][i]
		}
		for j, x := range nums {
			if s>>j&1 > 0 && (nums[i]%x == 0 || x%nums[i] == 0) {
				res += dfs(s^(1<<j), j)
			}
		}
		mem[s][i] = res
		return
	}
	for i := range nums {
		ans += dfs(u^(1<<i), i)
	}
	return ans % (1e9 + 7)
}
