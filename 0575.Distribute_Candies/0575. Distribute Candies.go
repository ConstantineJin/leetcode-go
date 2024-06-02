package main

func distributeCandies(candyType []int) int {
	mp := make(map[int]int)
	for _, candy := range candyType {
		mp[candy]++
	}
	return min(len(candyType)/2, len(mp))
}
