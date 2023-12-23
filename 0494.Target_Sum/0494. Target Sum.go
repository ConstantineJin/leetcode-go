package main

// 递归
//func findTargetSumWays(nums []int, target int) int {
//	// 添加+的数和为p，添加-的数和为所有元素的和s-p，则p-(s-p)=target，则p=(s+t)/2，转化为0-1背包
//	for _, num := range nums {
//		target += num
//	}
//	if target < 0 || target%2 == 1 {
//		return 0
//	}
//	target /= 2
//	var n = len(nums)
//	var mem = make([][]int, n)
//	for i := range mem {
//		mem[i] = make([]int, target+1)
//		for j := range mem[i] {
//			mem[i][j] = -1
//		}
//	}
//	var dfs func(int, int) int
//	dfs = func(i int, c int) (res int) { // 前i个物品，容量不超过c的方案数
//		if i < 0 {
//			if c == 0 {
//				return 1
//			}
//			return 0
//		}
//		var p = &mem[i][c]
//		if *p != -1 {
//			return *p
//		}
//		defer func() { *p = res }()
//		if c < nums[i] { // 当前物品比剩余容量大
//			return dfs(i-1, c) // 只能不选
//		}
//		return dfs(i-1, c) + dfs(i-1, c-nums[i]) // 选和不选
//	}
//	return dfs(n-1, target)
//}

// 递推
//func findTargetSumWays(nums []int, target int) int {
//	for _, num := range nums {
//		target += num
//	}
//	if target < 0 || target%2 == 1 {
//		return 0
//	}
//	target /= 2
//	var n = len(nums)
//	var f = make([][]int, n+1)
//	for i := range f {
//		f[i] = make([]int, target+1)
//	}
//	f[0][0] = 1
//	for i, num := range nums {
//		for c := 0; c <= target; c++ {
//			if c < num {
//				f[i+1][c] = f[i][c]
//			} else {
//				f[i+1][c] = f[i][c] + f[i][c-num]
//			}
//		}
//	}
//	return f[n][target]
//}

// 递推空间优化
func findTargetSumWays(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2
	var n = len(nums)
	var f = [2][]int{make([]int, target+1), make([]int, target+1)}
	f[0][0] = 1
	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if c < num {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = f[i%2][c] + f[i%2][c-num]
			}
		}
	}
	return f[n%2][target]
}
