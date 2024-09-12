package main

func countConsistentStrings(allowed string, words []string) (ans int) {
	set := make(map[rune]struct{})
	for _, c := range allowed {
		set[c] = struct{}{}
	}
next:
	for _, word := range words {
		for _, c := range word {
			if _, ok := set[c]; !ok {
				continue next
			}
		}
		ans++
	}
	return
}
