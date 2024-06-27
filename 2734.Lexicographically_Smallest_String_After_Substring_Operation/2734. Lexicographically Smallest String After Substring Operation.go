package main

func smallestString(str string) string {
	n := len(str)
	s := []byte(str)
	for i, c := range s {
		if c > 'a' {
			for j := i; j < n && s[j] > 'a'; j++ {
				s[j]--
			}
			return string(s)
		}
	}
	s[n-1] = 'z'
	return string(s)
}
