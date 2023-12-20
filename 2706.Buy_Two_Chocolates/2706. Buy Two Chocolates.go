package main

import "math"

func buyChoco(prices []int, money int) int {
	var a, b = math.MaxInt, math.MaxInt
	for _, price := range prices {
		a, b = min(a, price), min(b, max(a, price))
	}
	if a+b > money {
		return money
	} else {
		return money - a - b
	}
}
