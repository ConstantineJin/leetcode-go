package main

func canEat(candiesCount []int, queries [][]int) []bool {
	prefixSum := make([]int, len(candiesCount)+1)
	for i, count := range candiesCount {
		prefixSum[i+1] = prefixSum[i] + count
	}
	ans := make([]bool, len(queries))
	for i, query := range queries {
		//ans[i] = prefixSum[query[0]]/query[2] <= query[1] && query[1] < prefixSum[query[0]]+candiesCount[query[0]]
		favoriteType, favoriteDay, dailyCap := query[0], query[1], query[2]
		mn := prefixSum[favoriteType] / dailyCap                   // 每天吃 dailyCap 颗，吃到 favoriteType 的第一颗
		mx := prefixSum[favoriteType] + candiesCount[favoriteType] // 每天吃一颗，吃到 favoriteType 的最后一颗
		ans[i] = mn <= favoriteDay && favoriteDay < mx
	}
	return ans
}
