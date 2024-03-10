package main

import "strconv"

func getHint(secret string, guess string) string {
	var cntA, cntB int
	var cntSecret, cntGuess [10]int
	for i := range secret {
		if secret[i] == guess[i] {
			cntA++
		} else {
			cntSecret[secret[i]-'0']++
			cntGuess[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cntB += min(cntSecret[i], cntGuess[i])
	}
	return strconv.Itoa(cntA) + "A" + strconv.Itoa(cntB) + "B"
}
