package main

func maxScoreSightseeingPair(values []int) (ans int) {
	var mx int // j 左边的 values[i] + i 的最大值
	for j, value := range values {
		ans = max(ans, mx+value-j)
		mx = max(mx, value+j)
	}
	return
}
