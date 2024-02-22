package main

func findJudge(n int, trust [][]int) int {
	var s, t = make([]int, n), make([]int, n)
	for _, tr := range trust {
		s[tr[0]-1]++
		t[tr[1]-1]++
	}
	for i, v := range s {
		if v == 0 && t[i] == n-1 {
			return i + 1
		}
	}
	return -1
}
