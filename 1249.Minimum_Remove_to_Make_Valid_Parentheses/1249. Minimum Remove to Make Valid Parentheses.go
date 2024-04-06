package main

func minRemoveToMakeValid(s string) string {
	var depth int
	var ans []byte
	for i := range s {
		switch s[i] {
		case '(':
			ans = append(ans, s[i])
			depth++
		case ')':
			if depth > 0 {
				ans = append(ans, s[i])
				depth--
			}
		default:
			ans = append(ans, s[i])
		}
	}
	for i := len(ans) - 1; i >= 0 && depth > 0; i-- {
		if ans[i] == '(' {
			ans = append(ans[:i], ans[i+1:]...)
			depth--
		}
	}
	return string(ans)
}
