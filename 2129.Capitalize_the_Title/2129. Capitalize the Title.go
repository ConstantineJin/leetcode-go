package main

import "strings"

func capitalizeTitle(title string) (ans string) {
	var temp = strings.Split(title, " ")
	for _, s := range temp {
		if len(s) < 3 {
			ans += strings.ToLower(s)
		} else {
			ans += strings.Title(strings.ToLower(s))
		}
		ans += " "
	}
	return ans[:len(ans)-1]
}
