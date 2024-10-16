package main

func maxProfit(prices []int) (ans int) {
	mn := prices[0]
	for _, price := range prices[1:] {
		ans = max(ans, price-mn)
		mn = min(mn, price)
	}
	return
}
