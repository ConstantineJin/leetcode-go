package main

func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	color := make([]int, n)
	var isSafe func(int) bool
	isSafe = func(x int) bool {
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1
		for _, y := range graph[x] {
			if !isSafe(y) {
				return false
			}
		}
		color[x] = 2
		return true
	}
	for i := range n {
		if isSafe(i) {
			ans = append(ans, i)
		}
	}
	return
}
