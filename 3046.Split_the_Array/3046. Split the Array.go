package main

func isPossibleToSplit(nums []int) bool {
	mp := make(map[int]int)
	for _, num := range nums {
		if mp[num] == 2 {
			return false
		}
		mp[num]++
	}
	return true
}
