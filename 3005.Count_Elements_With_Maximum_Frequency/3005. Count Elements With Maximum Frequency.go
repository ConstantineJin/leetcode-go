package main

func maxFrequencyElements(nums []int) (ans int) {
	var mp = make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	var maxFreq int
	for _, v := range mp {
		if v > maxFreq {
			maxFreq, ans = v, v
		} else if v == maxFreq {
			ans += v
		}
	}
	return
}
