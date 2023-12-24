package main

import "math"

func minimumPerimeter(neededApples int64) int64 {
	var n = int64(math.Cbrt(float64(neededApples) / 4))
	if 2*n*(n+1)*(2*n+1) < neededApples {
		n++
	}
	return 8 * n
}
