package main

func wordSubsets(words1, words2 []string) (ans []string) {
	var maxCnt [26]int
	for _, word := range words2 {
		var cnt [26]int
		for _, c := range word {
			cnt[c-'a']++
		}
		for i := range 26 {
			maxCnt[i] = max(maxCnt[i], cnt[i])
		}
	}
next:
	for _, word := range words1 {
		var cnt [26]int
		for _, c := range word {
			cnt[c-'a']++
		}
		for i := range 26 {
			if cnt[i] < maxCnt[i] {
				continue next
			}
		}
		ans = append(ans, word)
	}
	return
}
