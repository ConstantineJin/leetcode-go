package main

func findMaximumScore(nums []int) (ans int64) {
	var mx int
	for _, num := range nums[:len(nums)-1] {
		mx = max(mx, num)
		ans += int64(mx)
	}
	return
}
