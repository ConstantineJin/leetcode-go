package main

import "sort"

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	var sunny []int            // 晴天的日期
	rainy := make(map[int]int) // 每个湖最近一次下雨的日期
	for i, rain := range rains {
		if rain == 0 { // 晴天
			sunny = append(sunny, i)
			ans[i] = 1 // 先随意抽干一个湖
		} else { // 雨天
			ans[i] = -1
			if j, ok := rainy[rain]; ok {
				idx := sort.Search(len(sunny), func(i int) bool { return sunny[i] > j }) // 二分查找该湖泊上一次下雨后的第一个晴天
				if idx == len(sunny) {                                                   // 不存在满足要求的晴天，发生洪水
					return nil
				}
				ans[sunny[idx]] = rain // 在那一天抽干当前湖的湖水
				sunny = append(sunny[:idx], sunny[idx+1:]...)
			}
			rainy[rain] = i // 更新当前湖最近一次下雨的时间
		}
	}
	return ans
}
