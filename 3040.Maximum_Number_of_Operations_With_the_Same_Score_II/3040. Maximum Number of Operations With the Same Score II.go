package main

func maxOperations(nums []int) (ans int) {
	n := len(nums)
	if n < 4 {
		return 1
	}
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	var clearMem func()
	clearMem = func() {
		for i := range mem {
			for j := range mem[i] {
				mem[i][j] = -1
			}
		}
	}
	var target int
	var dfs func(left, right int) (res int)
	dfs = func(left, right int) (res int) {
		if left >= right {
			return
		}
		if mem[left][right] > -1 {
			return mem[left][right]
		}
		if nums[left]+nums[right] == target {
			res = max(res, dfs(left+1, right-1)+1)
		}
		if nums[left]+nums[left+1] == target {
			res = max(res, dfs(left+2, right)+1)
		}
		if nums[right-1]+nums[right] == target {
			res = max(res, dfs(left, right-2)+1)
		}
		mem[left][right] = res
		return
	}
	target = nums[0] + nums[n-1]
	clearMem()
	ans = max(ans, dfs(1, n-2))
	target = nums[0] + nums[1]
	clearMem()
	ans = max(ans, dfs(2, n-1))
	target = nums[n-2] + nums[n-1]
	clearMem()
	ans = max(ans, dfs(0, n-3))
	return ans + 1
}
