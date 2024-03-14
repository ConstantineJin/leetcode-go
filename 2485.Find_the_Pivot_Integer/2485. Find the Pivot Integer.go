package main

import "math"

func pivotInteger(n int) (x int) {
	var temp = (n*n + n) / 2
	x = int(math.Sqrt(float64(temp)))
	if x*x == temp {
		return x
	}
	return -1
}
