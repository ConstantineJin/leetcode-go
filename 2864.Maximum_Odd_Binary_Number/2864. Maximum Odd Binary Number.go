package main

import "strings"

func maximumOddBinaryNumber(s string) string {
	var one = strings.Count(s, "1")
	return strings.Repeat("1", one-1) + strings.Repeat("0", len(s)-one) + "1"
}
