package main

func finalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{0} // 单调栈底存入 0 方便处理没有折扣的情况
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && prices[i] < st[len(st)-1] {
			st = st[:len(st)-1]
		}
		ans[i] = prices[i] - st[len(st)-1]
		st = append(st, prices[i])
	}
	return ans
}
