package main

// 方法1: 回溯
//func countMaxOrSubsets(nums []int) (ans int) {
//	var maxOr int
//	for _, num := range nums {
//		maxOr |= num
//	}
//	n := len(nums)
//	var dfs func(i, or int)
//	dfs = func(i, or int) {
//		if i == n {
//			if or == maxOr {
//				ans++
//			}
//			return
//		}
//		dfs(i+1, or)
//		dfs(i+1, or|nums[i])
//	}
//	dfs(0, 0)
//	return
//}

// 方法2: 枚举子集
func countMaxOrSubsets(nums []int) (ans int) {
	var maxOr int
	for _, num := range nums {
		maxOr |= num
	}
	or := make([]int, 1<<len(nums))
	for i, num := range nums {
		for mask, all := 0, 1<<i; mask < all; mask++ {
			res := or[mask] | num
			or[all|mask] = res
			if res == maxOr {
				ans++
			}
		}
	}
	return
}
