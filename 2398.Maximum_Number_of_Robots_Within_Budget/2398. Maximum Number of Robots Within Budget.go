package main

func maximumRobots(chargeTimes, runningCosts []int, _budget int64) (ans int) {
	budget := int(_budget)
	var queue []int // 单调递减队列
	var sum, left int
	for right, chargeTime := range chargeTimes { // 滑动窗口
		for len(queue) > 0 && chargeTime >= chargeTimes[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, right)
		sum += runningCosts[right]
		for len(queue) > 0 && chargeTimes[queue[0]]+(right-left+1)*sum > budget {
			if queue[0] == left { // 特别地，当队首就是滑动窗口左端点时，需要弹出队首
				queue = queue[1:]
			}
			sum -= runningCosts[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
