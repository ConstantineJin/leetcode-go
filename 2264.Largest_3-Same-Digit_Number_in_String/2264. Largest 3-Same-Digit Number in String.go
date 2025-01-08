package main

import "strings"

func largestGoodInteger(num string) string {
	mx, cnt := byte(0), 1
	for i := 1; i < len(num); i++ {
		d := num[i]
		if d != num[i-1] {
			cnt = 1
			continue
		}
		cnt++
		if cnt == 3 && d > mx {
			mx = d
		}
	}
	if mx == 0 {
		return ""
	}
	return strings.Repeat(string(mx), 3)
}
