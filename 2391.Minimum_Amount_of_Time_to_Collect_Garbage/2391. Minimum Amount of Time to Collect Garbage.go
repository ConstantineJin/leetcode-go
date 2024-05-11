package main

func garbageCollection(garbage []string, travel []int) int {
	ans, mp := len(garbage[0]), make(map[rune]int)
	var sumT int
	for i, g := range garbage[1:] {
		ans += len(g)
		sumT += travel[i]
		for _, c := range g {
			ans += sumT - mp[c]
			mp[c] = sumT
		}
	}
	return ans
}
