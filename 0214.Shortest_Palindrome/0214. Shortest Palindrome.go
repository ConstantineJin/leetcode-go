package main

import "slices"

func shortestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	temp := make([]byte, 1, n*2+1)
	temp[0] = '#'
	for i := range s {
		temp = append(temp, s[i])
		temp = append(temp, '#')
	}
	t := string(temp)
	m := len(t)
	armLength := make([]int, m)
	left, center := m, m
	for i := m - 1; i >= 0; i-- { // 倒序遍历 t，运行 Manacher 算法
		if i > left {
			armLength[i] = min(armLength[center*2-i], i-left)
		} else {
			armLength[i] = 1
		}
		for ; i-armLength[i] >= 0 && i+armLength[i] < m && t[i-armLength[i]] == t[i+armLength[i]]; armLength[i]++ {
		}
		if i-armLength[i] < left {
			left, center = i-armLength[i], i
		}
		if left < 1 { // 第一个覆盖到下标 0 的回文子串就是答案的回文中心
			reverse := []byte(s[armLength[i]-1:])
			slices.Reverse(reverse)
			return string(reverse) + s
		}
	}
	return ""
}
