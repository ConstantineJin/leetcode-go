package main

func trailingZeroes(n int) (ans int) {
	var five = 5
	for five <= n {
		ans += n / five
		five *= 5
	}
	return
}
