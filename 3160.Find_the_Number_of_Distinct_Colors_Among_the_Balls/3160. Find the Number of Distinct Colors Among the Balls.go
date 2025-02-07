package main

func queryResults(limit int, queries [][]int) []int {
	n := len(queries)
	color, cnt := make(map[int]int), make(map[int]int)
	ans := make([]int, n)
	for i, query := range queries {
		x, y := query[0], query[1]
		if c := color[x]; c > 0 {
			cnt[c]--
			if cnt[c] == 0 {
				delete(cnt, c)
			}
		}
		color[x] = y
		cnt[y]++
		ans[i] = len(cnt)
	}
	return ans
}
