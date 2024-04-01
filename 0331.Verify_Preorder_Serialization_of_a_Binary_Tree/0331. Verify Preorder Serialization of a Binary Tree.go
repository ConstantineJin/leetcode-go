package main

import "strings"

func isValidSerialization(preorder string) bool {
	var st []string
	for _, s := range strings.Split(preorder, ",") {
		st = append(st, s)
		for len(st) > 2 && st[len(st)-1] == "#" && st[len(st)-2] == "#" && st[len(st)-3] != "#" {
			st = st[:len(st)-2]
			st[len(st)-1] = "#"
		}
	}
	return len(st) == 1 && st[0] == "#"
}
