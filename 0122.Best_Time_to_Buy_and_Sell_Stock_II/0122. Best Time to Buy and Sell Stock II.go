package main

func maxProfit(prices []int) (ans int) {
	for i, price := range prices[1:] {
		ans += max(0, price-prices[i])
	}
	return
}
