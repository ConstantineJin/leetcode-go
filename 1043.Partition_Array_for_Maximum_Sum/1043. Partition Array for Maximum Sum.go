package main

// 递归
//func maxSumAfterPartitioning(arr []int, k int) int {
//	var n = len(arr)
//	var mem = make([]int, n)
//	for i := range mem {
//		mem[i] = -1
//	}
//	var dfs func(i int) int
//	dfs = func(i int) (res int) {
//		if i == n {
//			return 0
//		}
//		var p = &mem[i]
//		if *p != -1 {
//			return *p
//		}
//		defer func() { *p = res }()
//		for j := i + 1; j <= i+k && j <= n; j++ {
//			res = max(res, slices.Max(arr[i:j])*(j-i)+dfs(j))
//		}
//		return
//	}
//	return dfs(0)
//}

// 递推
func maxSumAfterPartitioning(arr []int, k int) int {
	var n = len(arr)
	var f = make([]int, k)
	for i := range arr {
		var res int
		for j, mx := i, 0; j > i-k && j >= 0; j-- {
			mx = max(mx, arr[j])
			res = max(res, f[j%k]+(i-j+1)*mx)
		}
		f[(i+1)%k] = res
	}
	return f[n%k]
}
