package main

func checkValidString(s string) bool {
	var leftSt, asteriskSt []int
	for i, ch := range s {
		switch ch {
		case '(':
			leftSt = append(leftSt, i)
		case '*':
			asteriskSt = append(asteriskSt, i)
		case ')':
			if len(leftSt) > 0 {
				leftSt = leftSt[:len(leftSt)-1]
			} else if len(asteriskSt) > 0 {
				asteriskSt = asteriskSt[:len(asteriskSt)-1]
			} else {
				return false
			}
		}
	}
	var i = len(leftSt) - 1
	for j := len(asteriskSt) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if leftSt[i] > asteriskSt[j] {
			return false
		}
	}
	return i < 0
}
