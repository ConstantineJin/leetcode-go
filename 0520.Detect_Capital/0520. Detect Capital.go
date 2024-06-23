package main

func detectCapitalUse(word string) bool {
	var capitalCnt int
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			capitalCnt++
		}
	}
	return capitalCnt == 0 || capitalCnt == len(word) || capitalCnt == 1 && word[0] >= 'A' && word[0] <= 'Z'
}
