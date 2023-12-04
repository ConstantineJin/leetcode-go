package main

// 递归，逆向
//func minCostClimbingStairs(cost []int) int {
//	n := len(cost)        // 台阶数
//	mem := make([]int, n) // 记忆化搜索，存储从各级台阶开始爬完的最小花费
//	for i := range mem {
//		mem[i] = -1 // 初始化mem中所有元素为-1
//	}
//	var dfs func(stair int) int // 从当前台阶开始爬完的最小花费
//	dfs = func(stair int) int {
//		if n-stair < 3 { // 如果当前台阶是最后两级中的一级
//			return 0
//		}
//		p := &mem[stair]
//		if *p != -1 { // 如果计算过从当前台阶开始爬完的最小花费
//			return *p // 直接返回对应值
//		}
//		*p = min(dfs(stair+1)+cost[stair+1], dfs(stair+2)+cost[stair+2]) // 动态规划，爬一级和爬两级两种方案中花费较小者；保存结果到mem中
//		return *p
//	}
//	return min(dfs(0)+cost[0], dfs(1)+cost[1]) // 可以选择从第0或者第1级开始爬
//}

// 递推，正向
func minCostClimbingStairs(cost []int) int {
	var pre, cur int
	for i := 2; i <= len(cost); i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}
