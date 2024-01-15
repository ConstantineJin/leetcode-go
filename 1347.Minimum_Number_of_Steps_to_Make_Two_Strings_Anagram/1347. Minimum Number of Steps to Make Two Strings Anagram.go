package main

func minSteps(s, t string) (ans int) {
	var cnt [26]int
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for _, ch := range t {
		cnt[ch-'a']--
	}
	for i := range cnt {
		ans += max(cnt[i], 0)
	}
	return
}
