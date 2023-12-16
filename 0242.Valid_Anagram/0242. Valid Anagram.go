package main

func isAnagram(s string, t string) bool {
	var mp = make(map[int32]int)
	for _, c := range s {
		mp[c]++
	}
	for _, c := range t {
		mp[c]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
