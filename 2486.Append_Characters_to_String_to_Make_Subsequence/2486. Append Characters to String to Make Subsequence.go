package main

func appendCharacters(s string, t string) int {
	var j int
	for i := 0; i < len(s) && j < len(t); i++ {
		if s[i] == t[j] {
			j++
		}
	}
	return len(t) - j
}
