package main

import "strings"

// 方法1: 中心扩展算法，时间复杂度 O(n^2)
//func expandAroundCenter(s string, left, right int) (int, int) {
//	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
//	}
//	return left + 1, right - 1
//}
//
//func longestPalindrome(s string) string {
//	if len(s) < 2 {
//		return s
//	}
//	var start, end int // 当前最长回文子串的始末下标
//	for i := 0; i < len(s); i++ {
//		left1, right1 := expandAroundCenter(s, i, i)   // 回文串长度为奇数
//		left2, right2 := expandAroundCenter(s, i, i+1) // 回文串长度为偶数
//		if right1-left1 > end-start {
//			start, end = left1, right1
//		}
//		if right2-left2 > end-start {
//			start, end = left2, right2
//		}
//	}
//	return s[start : end+1]
//}

// 方法2: Manacher 算法，时间复杂度 O(n)
func longestPalindrome(s string) string {
	// 预处理原字符串 s，在首尾及相邻字符间添加不在原串字符集中的字符，避免单独考虑长度为偶数的回文串
	t := make([]byte, 1, len(s)*2+1)
	t[0] = '#'
	for i := range s {
		t = append(t, s[i])
		t = append(t, '#')
	}
	s = string(t)

	armLength := make([]int, len(s)) // armLength 用于存储每个中心点的回文半径长度
	start, end := 0, -1              // start 和 end 用于记录最长回文子串的开始和结束位置
	right, j := -1, -1               // right 和 j 用于维护当前能扩展到的最右端和相应的中心点
	expand := func(left, right int) int {
		for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		}
		return (right-left)/2 - 1
	}
	for i := range s {
		var currentArmLength int
		if i <= right { // 如果当前下标没有超过右端点
			minArmLength := min(armLength[j*2-i], right-i)            // 利用关于 j 的对称性，以 i 为中心的最长回文子串半径等于其关于 j 的对称位置 j*2-i 的最长回文子串半径，但延伸不能超过 right
			currentArmLength = expand(i-minArmLength, i+minArmLength) // 基于朴素算法进行扩展
		} else { // 如果当前下标超过了右端点
			currentArmLength = expand(i, i) // 只能基于朴素算法进行扩展
		}
		armLength[i] = currentArmLength
		if i+currentArmLength > right { // 更新右端点及对应中心点
			j, right = i, i+currentArmLength
		}
		if currentArmLength*2+1 > end-start { // 更新最长回文子串的始末下标
			start, end = i-currentArmLength, i+currentArmLength
		}
	}
	return strings.ReplaceAll(s[start:end+1], "#", "") // 移除所有在预处理时添加的额外字
}
