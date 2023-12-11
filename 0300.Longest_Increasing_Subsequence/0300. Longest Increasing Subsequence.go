package main

func lengthOfLIS(nums []int) (ans int) {
	var n = len(nums)
	var mem = make([]int, n)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		var p = &mem[i]
		if *p > 0 {
			return *p
		}
		for j, v := range nums[:i] {
			if v < nums[i] {
				res = max(res, dfs(j))
			}
		}
		res++
		*p = res
		return
	}
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return
}
