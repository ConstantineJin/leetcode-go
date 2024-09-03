package main

import "strconv"

func getLucky(s string, k int) (ans int) {
	var t string
	for _, c := range s {
		t += strconv.Itoa(int(c-'a') + 1)
	}
	for i := 0; i < k; i++ {
		var sum int32
		for _, c := range t {
			sum += c - '0'
		}
		t = strconv.Itoa(int(sum))
	}
	ans, _ = strconv.Atoi(t)
	return
}
