package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := 0; i < len(nums)-indexDifference; i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
