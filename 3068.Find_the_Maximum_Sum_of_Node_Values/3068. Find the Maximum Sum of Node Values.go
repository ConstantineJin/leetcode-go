package main

import "math"

// 引理1:一个数异或2次（偶数次）k后保持不变
// 引理2:对于树上的从x到y的简单路径，如果对路径上的每条边都进行异或操作，则除x和y外路径上的所有节点都进行了偶数次异或操作，值保持不变。因此题目等价于可以对任意两个节点同时对k进行异或操作，也就不需要建树了
// 引理3:总有偶数个元素对k进行了异或操作
// 题目变为：任选nums中偶数个元素对k进行异或操作，使得nums各元素之和最大
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	// 定义 f[i][0], f[i][1]分别表示选择 nums 的前 i 个元素中偶数/奇数个元素异或 k 所能得到的最大和
	f0, f1 := 0, math.MinInt
	for _, num := range nums {
		f0, f1 = max(f0+num, f1+(num^k)), max(f1+num, f0+(num^k))
	}
	return int64(f0)
}
