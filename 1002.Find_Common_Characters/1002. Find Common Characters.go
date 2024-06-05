package main

import "math"

func commonChars(words []string) (ans []string) {
	var minFreq [26]int
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range words {
		var freq [26]int
		for _, c := range word {
			freq[c-'a']++
		}
		for i, f := range freq {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return
}
