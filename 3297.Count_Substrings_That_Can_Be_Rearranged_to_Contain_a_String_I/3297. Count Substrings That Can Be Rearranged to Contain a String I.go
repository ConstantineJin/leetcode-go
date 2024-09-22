package main

func validSubstringCount(word1, word2 string) (ans int64) {
	m, n := len(word1), len(word2)
	if m < n {
		return
	}
	var required, have [26]int
	for _, ch := range word2 {
		required[ch-'a']++
	}
	var left, missing int
	for _, cnt := range required {
		missing += cnt
	}
	for right, ch := range word1 {
		c := ch - 'a'
		have[c]++
		if have[c] <= required[c] {
			missing--
		}
		for missing == 0 {
			ans += int64(m - right)
			cLeft := word1[left] - 'a'
			if have[cLeft] <= required[cLeft] {
				missing++
			}
			have[cLeft]--
			left++
		}
	}
	return ans
}
