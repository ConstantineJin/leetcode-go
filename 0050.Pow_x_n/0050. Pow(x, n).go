package main

func myPow(x float64, n int) float64 {
	var quickMultiply func(i int) float64
	quickMultiply = func(i int) float64 {
		if i == 0 {
			return 1
		}
		var y = quickMultiply(i / 2)
		if i%2 == 0 {
			return y * y
		} else {
			return x * y * y
		}
	}
	if n >= 0 {
		return quickMultiply(n)
	} else {
		return 1 / quickMultiply(-n)
	}
}
