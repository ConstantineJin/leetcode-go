package main

func minCost(colors string, neededTime []int) (ans int) {
	var n = len(colors)
	var start int
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] {
			var mx int
			for _, time := range neededTime[start:i] {
				mx = max(mx, time)
				ans += time
			}
			ans -= mx
			start = i
		}
	}
	if colors[n-1] == colors[n-2] {
		var mx int
		for _, time := range neededTime[start:] {
			mx = max(mx, time)
			ans += time
		}
		ans -= mx
	}
	return
}
