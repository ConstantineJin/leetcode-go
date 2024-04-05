package main

func maxDepth(s string) (ans int) {
	var depth int
	for _, c := range s {
		switch c {
		case '(':
			depth++
			ans = max(ans, depth)
		case ')':
			depth--
		}
	}
	return
}
