package main

func reportSpam(message, bannedWords []string) bool {
	set := make(map[string]struct{}, len(bannedWords))
	for _, word := range bannedWords {
		set[word] = struct{}{}
	}
	var cnt int
	for _, word := range message {
		if _, ok := set[word]; ok {
			cnt++
		}
	}
	return cnt > 1
}
