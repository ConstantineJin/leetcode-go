package main

import "math"

// 递归
//func climbStairs(n int) int {
//	mem := make([]int, n+1) // 记忆化搜索
//	if n <= 2 {
//		return n
//	}
//	mem[1], mem[2] = 1, 2
//	var dfs func(k int) int
//	dfs = func(k int) int {
//		if mem[k] != 0 { // 如果当前剩余楼梯数的情况已经计算过了，依据mem中的值直接返回
//			return mem[k]
//		}
//		mem[k] = dfs(k-1) + dfs(k-2) // 否则需要计算，并保存结果
//		return mem[k]
//	}
//	return dfs(n)
//}

// 递推
//func climbStairs(n int) int {
//	prev, cur := 1, 1
//	for i := 2; i < n+1; i++ {
//		temp := cur
//		cur += prev
//		prev = temp
//	}
//	return cur
//}

// 通项公式
func climbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	return int(math.Round((math.Pow((1+sqrt5)/2, float64(n+1)) - math.Pow((1-sqrt5)/2, float64(n+1))) / sqrt5))
}
