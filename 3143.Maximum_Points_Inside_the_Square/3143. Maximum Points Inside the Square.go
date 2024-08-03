package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxPointsInsideSquare(points [][]int, s string) (ans int) {
	var min1 [26]int
	for i := range min1 {
		min1[i] = math.MaxInt
	}
	min2 := math.MaxInt
	for i, point := range points {
		d, tag := max(abs(point[0]), abs(point[1])), s[i]-'a'
		if d < min1[tag] {
			min2 = min(min2, min1[tag])
			min1[tag] = d
		} else if d < min2 {
			min2 = d
		}
	}
	for _, dis := range min1 {
		if dis < min2 {
			ans++
		}
	}
	return
}
