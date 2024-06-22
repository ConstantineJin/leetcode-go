package main

func numPairsDivisibleBy60(time []int) (ans int) {
	var remainder [60]int
	for _, t := range time {
		remainder[t%60]++
	}
	ans = (remainder[0]*(remainder[0]-1) + remainder[30]*(remainder[30]-1)) / 2
	for i := 1; i < 30; i++ {
		ans += remainder[i] * remainder[60-i]
	}
	return
}
