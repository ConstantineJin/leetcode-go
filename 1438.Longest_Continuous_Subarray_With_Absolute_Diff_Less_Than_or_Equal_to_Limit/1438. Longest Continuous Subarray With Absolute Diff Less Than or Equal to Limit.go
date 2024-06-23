package main

func longestSubarray(nums []int, limit int) (ans int) {
	var minQueue, maxQueue []int
	var left int
	for right, num := range nums {
		for len(minQueue) > 0 && minQueue[len(minQueue)-1] > num {
			minQueue = minQueue[:len(minQueue)-1]
		}
		minQueue = append(minQueue, num)
		for len(maxQueue) > 0 && maxQueue[len(maxQueue)-1] < num {
			maxQueue = maxQueue[:len(maxQueue)-1]
		}
		maxQueue = append(maxQueue, num)
		for len(minQueue) > 0 && len(maxQueue) > 0 && maxQueue[0]-minQueue[0] > limit {
			if nums[left] == minQueue[0] {
				minQueue = minQueue[1:]
			}
			if nums[left] == maxQueue[0] {
				maxQueue = maxQueue[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
