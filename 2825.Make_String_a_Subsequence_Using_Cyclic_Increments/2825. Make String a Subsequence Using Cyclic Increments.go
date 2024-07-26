package main

func canMakeSubsequence(str1, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	}
	var j int // 双指针
	for _, b := range str1 {
		c := byte(b) + 1
		if b == 'z' {
			c = 'a'
		}
		if byte(b) == str2[j] || c == str2[j] { // 子序列匹配，能匹配的就匹配
			j++
			if j == len(str2) {
				return true
			}
		}
	}
	return false
}
