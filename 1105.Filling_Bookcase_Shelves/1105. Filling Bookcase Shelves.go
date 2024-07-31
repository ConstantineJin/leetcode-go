package main

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	f := make([]int, n+1) // f[i+1] 表示把 books[0] 到 books[i] 按顺序摆放后的最小高度，递归边界 f[0]=0
	for i := range books {
		f[i+1] = math.MaxInt
		currentHeight, currentWidth := 0, shelfWidth
		for j := i; j >= 0; j-- {
			currentWidth -= books[j][0]
			if currentWidth < 0 { // 空间不足，无法放书
				break
			}
			currentHeight = max(currentHeight, books[j][1])
			f[i+1] = min(f[i+1], f[j]+currentHeight)
		}
	}
	return f[n]
}
