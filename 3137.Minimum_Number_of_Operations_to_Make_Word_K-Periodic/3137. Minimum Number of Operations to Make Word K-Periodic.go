package main

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	count := make(map[string]int)
	for i := k; i <= n; i += k {
		count[word[i-k:i]]++
	}
	var mx int
	for _, cnt := range count {
		mx = max(mx, cnt)
	}
	return n/k - mx
}
