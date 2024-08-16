package main

import (
	"math"
	"slices"
)

const inf = math.MaxInt >> 1

// 方法1: 动态规划，记忆化搜索
//func minimumValueSum(nums, andValues []int) int {
//	n, m := len(nums), len(andValues)
//	type args struct{ i, j, and int }
//	mem := make(map[args]int)
//	var dfs func(i, j, and int) (res int) // nums[i], andValues[j], 当前子数组 AND 结果为 and
//	dfs = func(i, j, and int) (res int) {
//		if n-i < m-j {
//			return inf
//		}
//		if j == m {
//			if i == n {
//				return
//			}
//			return inf
//		}
//		if v, ok := mem[args{i, j, and}]; ok {
//			return v
//		}
//		and &= nums[i]
//		if and < andValues[j] {
//			return inf
//		}
//		res = dfs(i+1, j, and)
//		if and == andValues[j] {
//			res = min(res, dfs(i+1, j+1, -1)+nums[i]) // -1 的补码为全 1
//		}
//		mem[args{i, j, and}] = res
//		return
//	}
//	if ans := dfs(0, 0, -1); ans != inf {
//		return ans
//	}
//	return -1
//}

// 方法2: 单调队列+logTrick+三指针
func minimumValueSum(nums, andValues []int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = inf
	}
	newF := make([]int, n+1)
	for _, target := range andValues {
		nums := slices.Clone(nums)
		left, right := 0, 0
		var q []int // 单调队列，保存 f 的下标
		var qi int  // 单调队列目前处理到 f[qi]

		newF[0] = inf
		for i, x := range nums {
			for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
				nums[j] &= x
			}
			for left <= i && nums[left] < target {
				left++
			}
			for right <= i && nums[right] <= target {
				right++
			}

			// 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值+1
			// 下面是单调队列的滑窗过程

			if left < right {
				// 单调队列：右边入
				for ; qi < right; qi++ {
					for len(q) > 0 && f[qi] <= f[q[len(q)-1]] {
						q = q[:len(q)-1]
					}
					q = append(q, qi)
				}

				// 单调队列：左边出
				for q[0] < left {
					q = q[1:]
				}

				// 单调队列：计算答案
				newF[i+1] = f[q[0]] + x // 队首就是最小值
			} else {
				newF[i+1] = inf
			}
		}
		f, newF = newF, f
	}
	if f[n] < inf {
		return f[n]
	}
	return -1
}
