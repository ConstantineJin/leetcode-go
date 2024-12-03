package main

func addSpaces(s string, spaces []int) string {
	spaces = append(spaces, len(s))
	ans := []byte(s[:spaces[0]])
	for i := 1; i < len(spaces); i++ {
		ans = append(append(ans, ' '), s[spaces[i-1]:spaces[i]]...)
	}
	return string(ans)
}
