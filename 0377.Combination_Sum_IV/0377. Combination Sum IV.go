package main

func combinationSum4(nums []int, target int) int {
	var mem = make([]int, target+1)
	for i := range mem {
		mem[i] = -1
	}
	var dfs func(i int) (res int)
	dfs = func(i int) (res int) {
		if i == 0 {
			return 1
		}
		var p = &mem[i]
		if *p > -1 {
			return *p
		}
		for _, num := range nums {
			if num <= i {
				res += dfs(i - num)
			}
		}
		*p = res
		return
	}
	return dfs(target)
}
