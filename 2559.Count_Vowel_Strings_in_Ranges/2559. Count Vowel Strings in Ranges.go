package main

var vowel = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	prefixSum := make([]int, n+1)
	for i, word := range words {
		if vowel[word[0]] && vowel[word[len(word)-1]] {
			prefixSum[i+1] = prefixSum[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i]
		}
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = prefixSum[query[1]+1] - prefixSum[query[0]]
	}
	return ans
}
