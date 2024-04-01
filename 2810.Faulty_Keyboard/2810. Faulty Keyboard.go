package main

import "slices"

func finalString(s string) string {
	var ss = []byte(s)
	for i := 0; i < len(ss); {
		if ss[i] == 'i' {
			tmp := ss[i+1:]
			ss = ss[:i]
			slices.Reverse(ss)
			ss = append(ss, tmp...)
		} else {
			i++
		}
	}
	return string(ss)
}
