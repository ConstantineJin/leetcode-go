package main

func imageSmoother(img [][]int) [][]int {
	var m, n = len(img), len(img[0])
	var ans = make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			var sum, cnt int
			for _, row := range img[max(i-1, 0):min(i+2, m)] {
				for _, v := range row[max(j-1, 0):min(j+2, n)] {
					sum += v
					cnt++
				}
			}
			ans[i][j] = sum / cnt
		}
	}
	return ans
}
