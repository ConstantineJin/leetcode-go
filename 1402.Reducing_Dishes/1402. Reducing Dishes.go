package main

import "sort"

func maxSatisfaction(satisfaction []int) (res int) {
	sort.Ints(satisfaction)                                                  //贪心算法，对菜品满意度升序排序
	var sum int                                                              //sum为做菜的满意度的非加权和
	for i := len(satisfaction) - 1; i >= 0 && sum+satisfaction[i] > 0; i-- { //只要新增的菜的满意度与之前的所有菜的非加权满意度之和（即增加一个时间单位带来的总和满意度提升）大于0
		sum += satisfaction[i]
		res += sum
	}
	return
}
