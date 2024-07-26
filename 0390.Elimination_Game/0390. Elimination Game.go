package main

func lastRemaining(n int) int {
	cnt := n             // 当前数组长度
	first := 1           // 当前数组从左往右的第一个元素
	step := 1            // 步长
	rightToLeft := false // 删除方向
	for cnt > 1 {
		if !(rightToLeft && cnt%2 == 0) { // 除非从右往左删且当前数组长度为偶数，否则 first 将变为 first+step
			first += step
		}
		rightToLeft = !rightToLeft // 反向
		step *= 2
		cnt /= 2
	}
	return first
}
