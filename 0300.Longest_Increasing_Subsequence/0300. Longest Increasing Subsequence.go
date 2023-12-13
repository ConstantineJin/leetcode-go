package main

import "sort"

// 递归
//func lengthOfLIS(nums []int) (ans int) {
//	var mem = make([]int, len(nums)) // 记忆化搜索，记录以nums[i]结尾的LIS长度，由于LIS长度至少是1（这个元素本身），因此mem数组可以用0表示未计算过
//	var dfs func(i int) int  // 返回以nums[i]结尾的LIS长度
//	dfs = func(i int) (res int) {
//		var p = &mem[i]
//		if *p > 0 { // 已经计算过了
//			return *p
//		}
//		for j, v := range nums[:i] { // 从头枚举到nums[i-1]
//			if v < nums[i] { // 如果nums[j]比nums[i]小
//				res = max(res, dfs(j)) // 递归
//			}
//		}
//		res++    // 加上nums[i]自己
//		*p = res // 保存搜索结果
//		return
//	}
//	for i := range nums {
//		ans = max(ans, dfs(i))
//	}
//	return
//}

// 递推
//func lengthOfLIS(nums []int) (ans int) {
//	var f = make([]int, len(nums))
//	for i := range nums {
//		for j := range nums[:i] {
//			if nums[j] < nums[i] {
//				f[i] = max(f[i], f[j])
//			}
//		}
//		f[i]++
//	}
//	for _, v := range f {
//		ans = max(ans, v)
//	}
//	return
//}

// 贪心+二分
//func lengthOfLIS(nums []int) int {
//	var g = make([]int, 0) // g[i]表示长度为i+1的IS的末尾元素的最小值
//	for _, num := range nums {
//		var j = sort.SearchInts(g, num)
//		if j == len(g) { // >=num 的g[j]不存在
//			g = append(g, num)
//		} else {
//			g[j] = num
//		}
//	}
//	return len(g)
//}

// 原地修改
func lengthOfLIS(nums []int) int {
	var g = nums[:0]
	for _, num := range nums {
		var j = sort.SearchInts(g, num)
		if j == len(g) { // >=num 的g[j]不存在
			g = append(g, num)
		} else {
			g[j] = num
		}
	}
	return len(g)
}
