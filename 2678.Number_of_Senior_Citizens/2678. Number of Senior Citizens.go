package main

import "strconv"

func countSeniors(details []string) (res int) {
	for _, detail := range details {
		age, _ := strconv.Atoi(detail[11:13])
		if age > 60 {
			res++
		}
	}
	return
}
