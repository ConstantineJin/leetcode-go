package main

// validStandardPalindrome 判断字符串s是否为标准回文串
func validStandardPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return validStandardPalindrome(s[i:j]) || validStandardPalindrome(s[i+1:j+1]) // 分别判断删除s[i]或s[j]后是否为标准回文串，只要有一个是就返回true
		}
		i++
		j--
	}
	return true
}
