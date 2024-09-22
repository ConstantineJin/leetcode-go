package main

func minWindow(s, t string) string {
	ansLeft, ansRight := -1, len(s)
	var left, less int
	var cntS, cntT [128]int
	for _, c := range t {
		if cntT[c] == 0 {
			less++
		}
		cntT[c]++
	}
	for right, c := range s {
		cntS[c]++
		if cntS[c] == cntT[c] {
			less--
		}
		for less == 0 {
			if right-left < ansRight-ansLeft {
				ansLeft, ansRight = left, right
			}
			out := s[left]
			if cntS[out] == cntT[out] {
				less++
			}
			cntS[out]--
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}
