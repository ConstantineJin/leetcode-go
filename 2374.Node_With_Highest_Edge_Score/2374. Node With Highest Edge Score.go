package main

func edgeScore(edges []int) (ans int) {
	scores := make([]int, len(edges))
	for i, edge := range edges {
		scores[edge] += i
		if scores[edge] > scores[ans] || scores[edge] == scores[ans] && edge < ans {
			ans = edge
		}
	}
	return
}
