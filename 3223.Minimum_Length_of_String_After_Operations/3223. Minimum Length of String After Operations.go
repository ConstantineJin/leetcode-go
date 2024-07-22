package main

func minimumLength(s string) (ans int) {
	var count [26]int
	for _, c := range s {
		count[c-'a']++
	}
	for _, cnt := range count {
		ans += (cnt-1)%2 + 1
	}
	return
}
