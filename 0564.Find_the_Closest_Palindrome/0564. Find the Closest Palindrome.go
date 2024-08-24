package main

import (
	"math"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nearestPalindromic(n string) string {
	m := len(n)
	candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1} // 备选项包括 100……001 和 999……999
	selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])                            // 原数的前半部份
	for _, x := range []int{selfPrefix - 1, selfPrefix, selfPrefix + 1} { // 备选项包括原数的前半部份 -1、保持不变和 +1 替换后半部份的结果
		y := x
		if m%2 == 1 {
			y /= 10
		}
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}
	ans := -1
	selfNumber, _ := strconv.Atoi(n)
	for _, candidate := range candidates {
		if candidate != selfNumber {
			if ans == -1 || abs(candidate-selfNumber) < abs(ans-selfNumber) || abs(candidate-selfNumber) == abs(ans-selfNumber) && candidate < ans {
				ans = candidate
			}
		}
	}
	return strconv.Itoa(ans)
}
