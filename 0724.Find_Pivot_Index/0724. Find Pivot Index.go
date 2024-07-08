package main

func pivotIndex(nums []int) int {
	var s, leftS int
	for _, num := range nums {
		s += num
	}
	for i, num := range nums {
		if leftS*2 == s-num {
			return i
		}
		leftS += num
	}
	return -1
}
