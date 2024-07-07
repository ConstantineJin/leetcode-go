package main

func numWaterBottles(numBottles int, numExchange int) (ans int) {
	empty, full := 0, numBottles
	for empty+full >= numExchange {
		ans += full
		empty += full
		full = empty / numExchange
		empty = empty % numExchange
	}
	return ans + full
}
