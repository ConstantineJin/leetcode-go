package main

func mySqrt(x int) (ans int) {
	var l, r = 0, x
	for l <= r {
		var m = (l + r) / 2
		if m*m <= x {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return
}
