package main

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	return sort.Search((position[len(position)-1]-position[0])/(m-1), func(distance int) bool {
		cnt, pre := 1, position[0] // 必定选择第一个篮子
		for i := 1; i < len(position) && cnt < m; i++ {
			if position[i]-pre > distance {
				cnt++
				pre = position[i]
			}
		}
		return cnt < m
	})
}
