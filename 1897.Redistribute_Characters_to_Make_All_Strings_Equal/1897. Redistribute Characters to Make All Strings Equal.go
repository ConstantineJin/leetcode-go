package main

func makeEqual(words []string) bool {
	var n, cnt = len(words), make([]int, 26)
	for _, word := range words {
		for i := range word {
			cnt[word[i]-'a']++
		}
	}
	for _, v := range cnt {
		if v%n != 0 {
			return false
		}
	}
	return true
}
