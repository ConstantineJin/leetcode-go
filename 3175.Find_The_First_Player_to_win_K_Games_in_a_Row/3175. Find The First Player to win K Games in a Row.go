package main

func findWinningPlayer(skills []int, k int) (ans int) {
	var win int
	for i := 1; i < len(skills) && win < k; i++ {
		if skills[i] > skills[ans] {
			ans = i
			win = 0
		}
		win++
	}
	return
}
