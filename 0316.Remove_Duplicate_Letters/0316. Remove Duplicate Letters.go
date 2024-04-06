package main

func removeDuplicateLetters(s string) string {
	var cnt, inAns = make(map[rune]int), make(map[rune]bool)
	for _, c := range s {
		cnt[c]++
	}
	var ans []rune
	for _, c := range s {
		cnt[c]--
		if inAns[c] {
			continue
		}
		for len(ans) > 0 && c < ans[len(ans)-1] && cnt[ans[len(ans)-1]] > 0 {
			inAns[ans[len(ans)-1]] = false
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, c)
		inAns[c] = true
	}
	return string(ans)
}
