package main

func sumOfTheDigitsOfHarshadNumber(x int) (ans int) {
	for y := x; y > 0; y /= 10 {
		ans += y % 10
	}
	if x%ans == 0 {
		return
	} else {
		return -1
	}
}
