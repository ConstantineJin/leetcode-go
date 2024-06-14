package main

import "sort"

const mx = 1e5 // nums[i]的取值上限

// 方法1:计数
//func minIncrementForUnique(nums []int) (ans int) {
//	count := make([]int, mx*2) // 最坏情况下，nums长为1e5且每个元素也都是1e5，需要变成1e5～2e5-1
//	for _, num := range nums {
//		count[num]++
//	}
//	var repeatCnt int // 重复出现的次数总和
//	for i, cnt := range count {
//		if cnt > 1 {
//			repeatCnt += cnt - 1
//			ans -= i * (cnt - 1)
//		} else if repeatCnt > 0 && cnt == 0 { // 贪心策略：当找到一个未出现过的数时，将之前某个重复出现的数增加成这个没有出现过的数
//			repeatCnt--
//			ans += i
//		}
//	}
//	return
//}

// 方法2:排序
func minIncrementForUnique(nums []int) (ans int) {
	sort.Ints(nums) // 非递减排序
	for i, num := range nums[:len(nums)-1] {
		if num >= nums[i+1] { // 如果出现不严格递增
			ans += num - nums[i+1] + 1
			nums[i+1] += num - nums[i+1] + 1
		}
	}
	return
}
