package main

import "slices"

func minimumCoins(prices []int) int {
	n := len(prices)
	for i := (n+1)/2 - 1; i > 0; i-- {
		prices[i-1] += slices.Min(prices[i : i*2+1]) // prices[i] 表示在购买第 i 个水果的前提下获得第 i 个及其后的水果所需的最少金币数
	}
	return prices[0]
}
