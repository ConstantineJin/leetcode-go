package main

func permute(nums []int) (ans [][]int) {
	var n = len(nums)
	var set, pmt = make([]bool, n), make([]int, n) // set记录每个元素是否在排列中出现过，pmt记录当前的排列
	var dfs func(int)
	dfs = func(i int) {
		if i == n { // 完成了一次排列
			ans = append(ans, append([]int(nil), pmt...)) // 添加到结果中
			return
		}
		for j, ok := range set {
			if !ok { // 当前元素还没有在排列中出现过
				pmt[i], set[j] = nums[j], true // 添加到排列中
				dfs(i + 1)                     // dfs
				set[j] = false                 // 从排列中移除当前元素
			}
		}
	}
	dfs(0)
	return
}
