package main

import "math"

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i + j*j
		if sum < c {
			i++
		} else if sum > c {
			j--
		} else {
			return true
		}

	}
	return false
}
