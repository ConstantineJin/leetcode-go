package main

func minOperations(nums []int) (ans int) {
	var mp = make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for _, v := range mp {
		if v == 1 {
			return -1
		}
		ans += (v + 2) / 3
	}
	return
}
