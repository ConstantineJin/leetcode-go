package main

import "strings"

func doesAliceWin(s string) bool {
	return strings.ContainsAny(s, "aeiou")
}
