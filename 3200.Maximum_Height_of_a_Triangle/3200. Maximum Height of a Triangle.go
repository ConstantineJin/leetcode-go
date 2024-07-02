package main

import "math"

// 方法1: 枚举
//func maxHeightOfTriangle(red int, blue int) int {
//	var cnt [2]int
//	for i := 1; ; i++ {
//		cnt[i%2] += i
//		if (cnt[0] > red || cnt[1] > blue) && (cnt[0] > blue || cnt[1] > red) {
//			return i - 1
//		}
//	}
//}

// 方法2：数学
func f(m, n int) int {
	odd := int(math.Sqrt(float64(m)))
	even := int((math.Sqrt(float64(n*4+1)) - 1) / 2)
	if odd > even {
		return even*2 + 1
	}
	return odd * 2
}

func maxHeightOfTriangle(red int, blue int) int {
	return max(f(red, blue), f(blue, red))
}
