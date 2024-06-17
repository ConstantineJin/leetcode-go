package main

func countCompleteDayPairs(hours []int) int64 {
	var ans int
	var arr [24]int
	for _, hour := range hours {
		arr[hour%24]++
	}
	for i := 1; i < 12; i++ {
		ans += arr[i] * arr[24-i]
	}
	ans += (arr[0]*(arr[0]-1) + arr[12]*(arr[12]-1)) / 2
	return int64(ans)
}
