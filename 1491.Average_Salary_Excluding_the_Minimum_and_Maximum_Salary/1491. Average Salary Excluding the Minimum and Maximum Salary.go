package main

func average(salary []int) float64 {
	var sum, mn, mx = salary[0], salary[0], salary[0]
	for _, s := range salary[1:] {
		sum += s
		mn, mx = min(mn, s), max(mx, s)
	}
	return float64(sum-mn-mx) / float64(len(salary)-2)
}
