package main

func canConstruct(ransomNote string, magazine string) bool {
	var cnt = [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		if cnt[ch-'a'] == 0 {
			return false
		}
		cnt[ch-'a']--
	}
	return true
}
