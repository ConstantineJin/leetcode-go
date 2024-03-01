package main

func twoSum(numbers []int, target int) []int {
	var left, right = 0, len(numbers) - 1
	for {
		var s = numbers[left] + numbers[right]
		if s == target {
			return []int{left + 1, right + 1}
		}
		if s > target {
			right--
		} else {
			left++
		}
	}
}
