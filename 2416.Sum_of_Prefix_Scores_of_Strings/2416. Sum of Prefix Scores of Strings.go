package main

type Trie struct {
	children [26]*Trie
	score    int
}

func sumPrefixScores(words []string) []int {
	root := &Trie{}
	for _, word := range words {
		cur := root
		for _, ch := range word {
			c := ch - 'a'
			if cur.children[c] == nil {
				cur.children[c] = &Trie{}
			}
			cur = cur.children[c]
			cur.score++
		}
	}
	ans := make([]int, len(words))
	for i, word := range words {
		cur := root
		for _, ch := range word {
			cur = cur.children[ch-'a']
			ans[i] += cur.score
		}
	}
	return ans
}
