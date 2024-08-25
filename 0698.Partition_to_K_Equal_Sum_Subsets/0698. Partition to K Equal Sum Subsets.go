package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k > 0 { // 如果整体的和不是 k 的倍数，直接返回 false
		return false
	}
	sum /= k // 每个子集的和
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > sum { // 如果最大元素比子集和还要大，直接返回 false
		return false
	}
	mem := make([]bool, 1<<n) // 用一个整数 s 来表示当前可用的数字集合：从低位到高位，第 i 位为 1 则表示数字 nums[i] 可以使用，否则表示 nums[i] 已被使用。mem[s] 来表示在可用的数字状态为 s 的情况下是否可行
	for i := range mem {
		mem[i] = true
	}
	var dfs func(s, p int) bool
	dfs = func(s, p int) bool {
		if s == 0 { // 如果没有剩余的数字，说明成功完成划分
			return true
		}
		if !mem[s] {
			return mem[s]
		}
		mem[s] = false
		for i, num := range nums {
			if p+num > sum {
				break
			}
			if s>>i&1 > 0 && dfs(s^1<<i, (p+num)%sum) { // 如果 nums[i] 尚未被使用，尝试将其加入子集。
				return true
			}
		}
		return false
	}
	return dfs(1<<n-1, 0)
}
