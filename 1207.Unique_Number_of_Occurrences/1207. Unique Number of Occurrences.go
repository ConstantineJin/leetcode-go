package main

func uniqueOccurrences(arr []int) bool {
	var occurrence, count = make(map[int]int), make(map[int]int)
	for _, value := range arr {
		occurrence[value]++
	}
	for _, value := range occurrence {
		if count[value] == 1 {
			return false
		} else {
			count[value] = 1
		}
	}
	return true
}
