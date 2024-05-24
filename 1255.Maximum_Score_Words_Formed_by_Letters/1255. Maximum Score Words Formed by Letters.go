package main

func maxScoreWords(words []string, letters []byte, score []int) (ans int) {
	alphabet := [26]int{}
	for _, letter := range letters {
		alphabet[letter-'a']++
	}
	n := len(words)
	requirement := make([][26]int, n)
	for i, word := range words {
		for _, c := range word {
			requirement[i][c-'a']++
		}
	}
	var dfs func(i, sc int)
	dfs = func(i, sc int) {
		if i == n {
			ans = max(ans, sc)
			return
		}
		flag := true
		for _, c := range words[i] {
			if requirement[i][c-'a'] > alphabet[c-'a'] {
				flag = false
				break
			}
		}
		dfs(i+1, sc)
		if flag {
			dfs(i+1, sc)
			for c, cnt := range requirement[i] {
				alphabet[c] -= cnt
				sc += cnt * score[c]
			}
			dfs(i+1, sc)
			for c, cnt := range requirement[i] {
				alphabet[c] += cnt
				sc -= cnt * score[c]
			}
		}
	}
	dfs(0, 0)
	return
}
