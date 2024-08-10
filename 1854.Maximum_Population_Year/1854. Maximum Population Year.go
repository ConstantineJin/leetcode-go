package main

func maximumPopulation(logs [][]int) (ans int) {
	var diff [101]int
	for _, log := range logs {
		diff[log[0]-1950]++
		diff[log[1]-1950]--
	}
	var mx, cnt int
	for i, d := range diff {
		cnt += d
		if cnt > mx {
			mx, ans = cnt, 1950+i
		}
	}
	return
}
