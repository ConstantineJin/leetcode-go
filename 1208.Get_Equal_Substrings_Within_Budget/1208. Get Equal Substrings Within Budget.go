package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func equalSubstring(s string, t string, maxCost int) (ans int) {
	var left, cost int
	for right := 0; right < len(s); right++ {
		cost += abs(int(s[right]) - int(t[right]))
		for ; cost > maxCost; left++ {
			cost -= abs(int(s[left]) - int(t[left]))
		}
		ans = max(ans, right-left+1)
	}
	return
}
