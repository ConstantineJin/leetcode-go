package main

// 方法1: 模拟
//func numWaterBottles(numBottles int, numExchange int) (ans int) {
//	empty, full := 0, numBottles
//	for empty+full >= numExchange {
//		ans += full
//		empty += full
//		full = empty / numExchange
//		empty = empty % numExchange
//	}
//	return ans + full
//}

// 方法2: 数学
func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return numBottles
	}
	return (numBottles-numExchange)/(numExchange-1) + 1 + numBottles
}
