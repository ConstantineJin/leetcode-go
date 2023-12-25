package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	var n, mp, ans = len(grid), make(map[int]bool), make([]int, 2)
	for _, r := range grid {
		for _, v := range r {
			if mp[v] == true {
				ans[0] = v
				continue
			}
			mp[v] = true
		}
	}
	for i := 1; i <= n*n; i++ {
		if _, ok := mp[i]; !ok {
			ans[1] = i
		}
	}
	return ans
}
