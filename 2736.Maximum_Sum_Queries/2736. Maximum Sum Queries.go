package main

import "sort"

// 单调栈+二分查找
func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	ans, sortedNums, sortedQueries, stack := make([]int, len(queries)), make([][]int, len(nums1)), make([][]int, len(queries)), make([][]int, 0)
	for i := range ans {
		ans[i] = -1 // 初始化ans全部元素为-1
	}
	for i := range nums1 {
		sortedNums[i] = []int{nums1[i], nums2[i]}
	}
	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i][0] > sortedNums[j][0] // sortedNums根据nums1进行降序排列
	})
	for i, query := range queries {
		sortedQueries[i] = []int{i, query[0], query[1]}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][1] > sortedQueries[j][1] // sortedQueries根据x进行降序排列
	})
	for _, query := range sortedQueries {
		for j := 0; j < len(sortedNums) && sortedNums[j][0] >= query[1]; j++ { // 保证nums1>=x
			num1, num2 := sortedNums[j][0], sortedNums[j][1]
			for len(stack) > 0 && stack[len(stack)-1][1] <= num1+num2 {
				stack = stack[:len(stack)-1] // 维护单调栈stack，每访问一个sortedNums[j]就将栈尾的具有更小的目标值的和的元素弹出
			}
			if len(stack) == 0 || stack[len(stack)-1][0] < num2 {
				stack = append(stack, []int{num2, num1 + num2})
			}
		}
		k := sort.Search(len(stack), func(i int) bool {
			return stack[i][0] >= query[2]
		})
		if k < len(stack) {
			ans[query[0]] = stack[k][1]
		}
	}
	return ans
}
