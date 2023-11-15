package main

import "strings"

// 直接使用双指针
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	str, i, j := []byte(s), 0, len(s)-1
	for i < j {
		//if str[i] >= 'A' && str[i] <= 'Z' {
		//	str[i] += 32
		//}
		//if str[j] >= 'A' && str[j] <= 'Z' {
		//	str[j] += 32
		//}
		if str[i] < '0' || str[i] > '9' && str[i] < 'a' || str[i] > 'z' {
			i++
			continue
		}
		if str[j] < '0' || str[j] > '9' && str[j] < 'a' || str[j] > 'z' {
			j--
			continue
		}
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
