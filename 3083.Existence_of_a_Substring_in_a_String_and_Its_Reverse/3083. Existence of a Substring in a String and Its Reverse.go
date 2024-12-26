package main

func isSubstringPresent(s string) bool {
	set := make(map[[2]byte]struct{})
	for i := range s[1:] {
		set[[2]byte{s[i], s[i+1]}] = struct{}{}
	}
	for i := len(s) - 1; i > 0; i-- {
		if _, ok := set[[2]byte{s[i], s[i-1]}]; ok {
			return true
		}
	}
	return false
}
