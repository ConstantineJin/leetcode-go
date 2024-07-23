package main

import (
	"math"
	"sort"
)

const mod int = 1e9 + 7

func sumOfPowers(nums []int, k int) int {
	sort.Ints(nums) // 子序列与原数组顺序无关，可以排序
	n := len(nums)
	mem := make(map[int]int)          // 记忆化搜索
	var dfs func(i, j, k, mn int) int // 当前处理到第 i 个元素，上一个选取的是第 j 个元素，还需要选取 k 个元素，当前最小差值为 mn
	dfs = func(i, j, k, mn int) int {
		if i == n {
			if k == 0 {
				return mn
			}
			return 0
		}
		if n-i < k { // 剩余元素数量不足
			return 0
		}
		key := mn<<18 | (i << 12) | (j << 6) | k
		if v, ok := mem[key]; ok {
			return v
		}
		ans := dfs(i+1, j, k, mn) // 不选当前元素
		if j == n {
			ans += dfs(i+1, i, k-1, mn)
		} else {
			ans += dfs(i+1, i, k-1, min(mn, nums[i]-nums[j]))
		}
		ans %= mod
		mem[key] = ans
		return ans
	}
	return dfs(0, n, k, math.MaxInt)
}
