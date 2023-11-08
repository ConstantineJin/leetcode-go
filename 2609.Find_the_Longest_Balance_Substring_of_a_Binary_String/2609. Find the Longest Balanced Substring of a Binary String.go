package main

//func findTheLongestBalancedSubstring(s string) (res int) {
//	var i, cnt0, cnt1 int
//	for i < len(s) {
//		for i < len(s) && s[i] == '1' {
//			i++
//		}
//		for i < len(s) && s[i] == '0' {
//			cnt0++
//			i++
//		}
//		for i < len(s) && s[i] == '1' && cnt1 < cnt0 {
//			cnt1++
//			i++
//		}
//		res = max(res, cnt1)
//		cnt0, cnt1 = 0, 0
//	}
//	return res * 2
//}

func findTheLongestBalancedSubstring(s string) (res int) {
	pre, cur := 0, 0 // pre和cur分别记录前一个和当前连续相同字符串的长度
	for i, c := range s {
		cur++
		if i == len(s)-1 || byte(c) != s[i+1] { // 如果当前连续相同字符串在此终结
			if c == '1' {
				res = max(res, min(pre, cur)) // 如果当前字符为1，则先选择pre和cur的较小值，然后选择其与先前最大值的较大值
			}
			pre, cur = cur, 0 // pre变成cur，cur重置为0
		}
	}
	return res * 2
}
