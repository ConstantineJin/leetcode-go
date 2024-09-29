package main

import "strings"

func countOfSubstrings(word string, k int) int {
	return f(word, k) - f(word, k+1) // 将恰好包含 k 个辅音字母的子字符串总数转化为包含至少 k 个辅音字母的子字符串总数减去至少包含 k+1 个辅音字母的子字符串总数
}

func f(word string, k int) (ans int) {
	vowels := make(map[rune]int)
	var consonant, left int
	for _, c := range word {
		if strings.ContainsRune("aeiou", c) {
			vowels[c]++
		} else {
			consonant++
		}
		for len(vowels) == 5 && consonant >= k {
			out := rune(word[left])
			if strings.ContainsRune("aeiou", out) {
				vowels[out]--
				if vowels[out] == 0 {
					delete(vowels, out)
				}
			} else {
				consonant--
			}
			left++
		}
		ans += left
	}
	return
}
