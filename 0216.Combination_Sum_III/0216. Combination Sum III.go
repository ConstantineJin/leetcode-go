package main

func combinationSum3(k int, n int) (ans [][]int) {
	var dfs func([]int, int)
	dfs = func(arr []int, s int) {
		if len(arr) == k-1 {
			if n-s > arr[len(arr)-1] && n-s < 10 {
				arr = append(arr, n-s)
				var temp = make([]int, k)
				copy(temp, arr)
				ans = append(ans, temp)
			}
			return
		}
		for j := arr[len(arr)-1] + 1; j < 10; j++ {
			if s+j >= n {
				return
			}
			arr = append(arr, j)
			dfs(arr, s+j)
			arr = arr[:len(arr)-1]
		}
	}
	for i := 1; i < min(n+1, 10); i++ {
		dfs([]int{i}, i)
	}
	return
}
