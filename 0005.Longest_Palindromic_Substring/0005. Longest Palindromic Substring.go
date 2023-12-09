package main

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0 // 当前最长回文子串的始末下标
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)   // 回文串长度为奇数
		left2, right2 := expandAroundCenter(s, i, i+1) // 回文串长度为偶数
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}
