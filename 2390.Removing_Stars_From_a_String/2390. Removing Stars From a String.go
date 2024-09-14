package main

func removeStars(s string) string {
	var st []rune
	for _, c := range s {
		if c == '*' {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}
