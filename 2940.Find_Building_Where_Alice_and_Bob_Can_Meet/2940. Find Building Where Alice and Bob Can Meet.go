package main

import "sort"

type pair struct{ height, index int }

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	qs := make([][]pair, len(heights))
	for i, query := range queries {
		a, b := query[0], query[1]
		if a > b {
			a, b = b, a // 保证 Bob 始终不在左侧
		}
		if a == b || heights[a] < heights[b] { // 如果两人在同一位置或者 Alice 所在建筑比 Bob 矮，那么 Alice 可以直接跳到 Bob 所在位置
			ans[i] = b
		} else {
			qs[b] = append(qs[b], pair{heights[a], i}) // 离线询问
		}
	}
	var stack []int // 单调栈，越靠近栈顶的位置对应的高度越小
	for i := len(heights) - 1; i >= 0; i-- {
		for _, q := range qs[i] {
			j := sort.Search(len(stack), func(i int) bool { return heights[stack[i]] <= q.height }) - 1 // 二分查找 Bob 右侧第一个比 Alice 所在建筑高的建筑
			if j == -1 {                                                                                // 不存在满足条件的建筑
				ans[q.index] = -1
			} else {
				ans[q.index] = stack[j]
			}
		}
		for len(stack) > 0 && heights[i] >= heights[stack[len(stack)-1]] { // 更新单调栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
