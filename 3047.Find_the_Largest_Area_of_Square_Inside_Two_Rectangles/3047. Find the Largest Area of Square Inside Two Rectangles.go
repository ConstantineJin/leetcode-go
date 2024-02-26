package main

import "sort"

type rectangle struct{ bottomLeftX, bottomLeftY, topRightX, topRightY int }

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	var n, ans = len(bottomLeft), 0
	var rectangles = make([]rectangle, n)
	for i := range rectangles {
		rectangles[i] = rectangle{bottomLeftX: bottomLeft[i][0], bottomLeftY: bottomLeft[i][1], topRightX: topRight[i][0], topRightY: topRight[i][1]}
	}
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i].bottomLeftX < rectangles[j].bottomLeftX })
	for i, r := range rectangles {
		for j := i + 1; j < n; j++ {
			var s = rectangles[j]
			if s.bottomLeftX >= r.topRightX {
				break
			} else if s.bottomLeftY > r.topRightY || s.topRightY < r.bottomLeftY {
				continue
			}
			var a, b = min(r.topRightX, s.topRightX) - max(r.bottomLeftX, s.bottomLeftX), min(r.topRightY, s.topRightY) - max(r.bottomLeftY, s.bottomLeftY)
			if a <= 0 || b <= 0 {
				continue
			} else {
				ans = max(ans, min(a, b))
			}
		}
	}
	return int64(ans * ans)
}
