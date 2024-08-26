package main

import "sort"

func answerQueries(nums, queries []int) []int {
	sort.Ints(nums) // 要求的是子序列，所以排序后不影响结果
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1] // 将 nums 改造成前缀和
	}
	for i, query := range queries {
		queries[i] = sort.SearchInts(nums, query+1) // 二分查找。复用 queries 存储答案
	}
	return queries
}
