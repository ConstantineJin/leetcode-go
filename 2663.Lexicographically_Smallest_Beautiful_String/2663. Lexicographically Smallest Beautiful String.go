package main

func smallestBeautifulString(str string, k int) string {
	limit := 'a' + byte(k)
	s := []byte(str)
	n := len(s)
	i := n - 1 // 贪心，优先考虑增加靠后的字母
	s[i]++     // 先加一
	for i < n {
		if s[i] == limit { // 需要进位
			if i == 0 { // 无法继续进位
				return ""
			}
			s[i] = 'a'
			i--
			s[i]++
		} else if i > 0 && s[i] == s[i-1] || i > 1 && s[i] == s[i-2] { // 存在回文子串
			s[i]++
		} else { // 修改前面的字符时可能导致后面出现新的回文子串
			i++
		}
	}
	return string(s)
}
