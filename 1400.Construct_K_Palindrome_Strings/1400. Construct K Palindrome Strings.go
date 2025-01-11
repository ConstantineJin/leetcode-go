package main

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	var cnt [26]int
	for _, c := range s {
		cnt[c-'a']++
	}
	var odd int
	for _, c := range cnt {
		if c&1 > 0 {
			odd++
		}
	}
	return odd <= k
}
