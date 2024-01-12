package main

func countWords(words1 []string, words2 []string) (ans int) {
	var cnt1, cnt2 = make(map[string]int), make(map[string]int)
	for _, word := range words1 {
		cnt1[word]++
	}
	for _, word := range words2 {
		cnt2[word]++
	}
	for word, cnt := range cnt1 {
		if cnt == 1 && cnt2[word] == 1 {
			ans++
		}
	}
	return
}
