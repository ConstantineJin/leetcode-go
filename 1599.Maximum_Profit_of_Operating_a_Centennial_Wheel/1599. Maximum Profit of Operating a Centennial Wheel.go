package main

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	var i = 1 // 轮转的次数
	var profit, maxProfit, wait int
	var ans = -1 // (利润最大时)轮转的次数，即题目的返回值，其初值为-1
	for ; wait > 0 || i <= len(customers); i++ {
		if i <= len(customers) {
			wait += customers[i-1] // 又来了些顾客，先加入等待队列
		}
		var board = min(4, wait)                   // 上摩天轮的人数
		wait -= board                              // 更新未上摩天轮的人数
		profit += board*boardingCost - runningCost // 更新利润
		if profit > maxProfit {
			maxProfit = profit
			ans = i
		}
	}
	return ans
}
