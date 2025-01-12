package main

func canBeValid(s, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var x int
	for i, c := range s {
		if c == '(' || locked[i] == '0' {
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}
	x = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}
	return true
}
