package main

// 暴力
//func minLength(s string) int {
//	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
//		s = strings.ReplaceAll(s, "AB", "")
//		s = strings.ReplaceAll(s, "CD", "")
//	}
//	return len(s)
//}

// 栈
func minLength(s string) int {
	var stack []byte
	for i := 0; i < len(s); i++ { // 遍历s
		if len(stack) > 0 && (s[i] == 'B' && stack[len(stack)-1] == 'A' || s[i] == 'D' && stack[len(stack)-1] == 'C') { // 如果当前字符是B或者D且对应的栈顶字符是A或者C
			stack = stack[:len(stack)-1] // 栈顶元素出栈
		} else {
			stack = append(stack, s[i]) // 否则当前字符入栈
		}
	}
	return len(stack) // 遍历完成后栈中元素数量即为答案
}
