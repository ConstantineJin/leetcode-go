package main

func secondHighest(s string) int {
	first, second := -1, -1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			v := int(c - '0')
			if v > first {
				first, second = v, first
			} else if v < first && v > second {
				second = v
			}
		}
	}
	return second
}
