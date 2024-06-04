package main

func minimumChairs(s string) (ans int) {
	var cur int
	for _, c := range s {
		if c == 'E' {
			cur++
			ans = max(ans, cur)
		} else {
			cur--
		}
	}
	return ans
}
