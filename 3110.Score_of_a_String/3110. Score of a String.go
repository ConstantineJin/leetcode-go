package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func scoreOfString(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		ans += abs(int(s[i]) - int(s[i-1]))
	}
	return
}
