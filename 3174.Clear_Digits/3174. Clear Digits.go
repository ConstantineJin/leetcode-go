package main

func clearDigits(s string) string {
	var ans []byte
	for i, c := range s {
		if '0' <= c && c <= '9' && len(ans) > 0 {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
