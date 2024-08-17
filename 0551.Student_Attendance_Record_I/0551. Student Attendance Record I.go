package main

func checkRecord(s string) bool {
	var a bool
	for i, c := range s {
		switch c {
		case 'A':
			if a {
				return false
			}
			a = true
		case 'L':
			if i > 1 && s[i-1] == 'L' && s[i-2] == 'L' {
				return false
			}
		}
	}
	return true
}
