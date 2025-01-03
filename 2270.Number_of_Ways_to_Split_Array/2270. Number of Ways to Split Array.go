package main

func waysToSplitArray(nums []int) (ans int) {
	var sum, leftSum int
	for _, num := range nums {
		sum += num
	}
	for _, num := range nums[:len(nums)-1] {
		leftSum += num
		if leftSum*2 >= sum {
			ans++
		}
	}
	return
}
