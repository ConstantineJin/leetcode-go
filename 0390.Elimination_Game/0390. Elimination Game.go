package main

func lastRemaining(n int) int {
	first := 1          // 当前数组从左往右的第一个元素
	step := 1           // 步长
	leftToRight := true // 删除方向
	for n > 1 {
		if leftToRight || n%2 == 1 { // 除非从右往左删且当前数组长度为偶数，否则 first 将变为 first+step
			first += step
		}
		leftToRight = !leftToRight // 反向
		step *= 2
		n /= 2
	}
	return first
}
