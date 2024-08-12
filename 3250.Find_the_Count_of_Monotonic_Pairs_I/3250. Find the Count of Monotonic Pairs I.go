package main

const mod = 1e9 + 7

func countOfPairs(nums []int) int {
	n := len(nums)
	mem := make([]map[int]int, n)
	for i := range mem {
		mem[i] = make(map[int]int)
	}
	var dfs func(i, prev int) (res int)
	dfs = func(i, prev int) (res int) {
		if i == n {
			return 1
		}
		if v, ok := mem[i][prev]; ok {
			return v
		}
		for j := prev; j <= nums[i]; j++ {
			if i == 0 || nums[i-1]-prev >= nums[i]-j {
				res = (res + dfs(i+1, j)) % mod
			}
		}
		mem[i][prev] = res
		return
	}
	return dfs(0, 0) % mod
}
