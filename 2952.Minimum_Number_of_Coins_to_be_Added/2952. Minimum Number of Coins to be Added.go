package main

import "sort"

func minimumAddedCoins(coins []int, target int) (ans int) {
	sort.Ints(coins)
	var s, i = 1, 0 // s代表得到了[0,s)区间内的所有整数
	for s <= target {
		if i < len(coins) && coins[i] <= s { // coins[i]不超过s
			s += coins[i] // 可以合并[0,s)和[coins[i],s+coins[i])得到[0,s+coins[i])
			i++
		} else { // 遍历完所有已有硬币，或当前遍历的硬币面额超过了s，就必须添加一枚面值为s的硬币
			s *= 2 // 得到了[0,2s)区间内的所有面值
			ans++
		}
	}
	return
}
