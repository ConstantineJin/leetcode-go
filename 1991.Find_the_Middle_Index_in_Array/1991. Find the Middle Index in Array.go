package main

func findMiddleIndex(nums []int) int {
	var sum, leftSum int
	for _, num := range nums {
		sum += num
	}
	for i, num := range nums {
		if leftSum*2 == sum-num {
			return i
		}
		leftSum += num
	}
	return -1
}
