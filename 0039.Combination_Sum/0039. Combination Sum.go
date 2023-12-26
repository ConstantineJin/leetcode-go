package main

func combinationSum(candidates []int, target int) (ans [][]int) {
	var comb, n = make([]int, 0), len(candidates)
	var dfs func(int, int)
	dfs = func(i, t int) { // 下标不小于i，目标和为t
		if i == n {
			return
		}
		if t == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		dfs(i+1, t)             // 不选当前元素
		if candidates[i] <= t { // 选当前元素的必要前提：当前元素不能超过剩余的目标和
			comb = append(comb, candidates[i]) // 选当前元素
			dfs(i, t-candidates[i])
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, target)
	return
}
