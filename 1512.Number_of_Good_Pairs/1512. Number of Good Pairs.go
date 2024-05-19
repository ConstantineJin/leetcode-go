package main

func numIdenticalPairs(nums []int) (ans int) {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for _, v := range mp {
		ans += v * (v - 1) / 2
	}
	return
}
