package main

func firstDayBeenInAllRooms(nextVisit []int) int {
	const mod int = 1e9 + 7
	var n = len(nextVisit)
	var s = make([]int, n)
	for i, v := range nextVisit[:n-1] {
		s[i+1] = (s[i]*2 - s[v] + 2 + mod) % mod
	}
	return s[n-1]
}
