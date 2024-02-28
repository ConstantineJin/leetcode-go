package main

func minIncrements(n int, cost []int) (ans int) {
	cost = append([]int{0}, cost...)
	for i := n / 2; i > 0; i-- {
		if cost[i*2] < cost[i*2+1] {
			ans += cost[i*2+1] - cost[i*2]
			cost[i] += cost[i*2+1]
		} else {
			ans += cost[i*2] - cost[i*2+1]
			cost[i] += cost[i*2]
		}
	}
	return
}
