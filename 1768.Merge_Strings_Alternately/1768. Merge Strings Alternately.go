package main

func mergeAlternately(word1, word2 string) string {
	m, n := len(word1), len(word2)
	ans := make([]byte, 0, m+n)
	for i := 0; i < max(m, n); i++ {
		if i < m {
			ans = append(ans, word1[i])
		}
		if i < n {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}
