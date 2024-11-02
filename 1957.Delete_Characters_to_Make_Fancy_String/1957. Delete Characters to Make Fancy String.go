package main

func makeFancyString(s string) string {
	var ans []byte
	for i := range s {
		if len(ans) > 1 && s[i] == ans[len(ans)-1] && s[i] == ans[len(ans)-2] {
			continue
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}
