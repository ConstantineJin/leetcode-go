package main

func maximumLength(nums []int) (ans int) {
	n := len(nums)
	var odd, even int
	for _, num := range nums {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	ans = max(odd, even)
	mem := make([][2]int, n)
	for i := range mem {
		mem[i] = [2]int{-1, -1}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == n {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		defer func() { mem[i][j] = res }()
		if nums[i]%2 == j {
			return max(dfs(i+1, j^1)+1, dfs(i+1, j))
		} else {
			return dfs(i+1, j)
		}
	}
	return max(ans, dfs(0, 0), dfs(0, 1))
}
