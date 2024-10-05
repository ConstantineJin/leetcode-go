package main

func checkInclusion(s1, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	var cnt1, cnt2 [26]int
	for _, c := range s1 {
		cnt1[c-'a']++
	}
	for _, c := range s2[:m] {
		cnt2[c-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i, c := range s2[m:] {
		cnt2[c-'a']++
		cnt2[s2[i]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
