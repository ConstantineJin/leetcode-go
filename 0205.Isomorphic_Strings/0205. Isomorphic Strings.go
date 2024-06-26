package main

func isIsomorphic(s string, t string) bool {
	var s2t, t2s = make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		if s2t[s[i]] > 0 && s2t[s[i]] != t[i] || t2s[t[i]] > 0 && t2s[t[i]] != s[i] {
			return false
		}
		s2t[s[i]], t2s[t[i]] = t[i], s[i]
	}
	return true
}
