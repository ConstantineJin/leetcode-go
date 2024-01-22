package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countOfPairs(n int, x int, y int) []int {
	var ans = make([]int, n)
	x, y = min(x, y)-1, max(x, y)-1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			ans[min(i-j-1, abs(x-j)+abs(y-i))] += 2
		}
	}
	return ans
}
