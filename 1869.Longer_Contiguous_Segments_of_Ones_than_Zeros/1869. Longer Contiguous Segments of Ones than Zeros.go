package main

func checkZeroOnes(s string) bool {
	n, i, j, cnt := len(s), 0, 0, make([]int, 2)
	for i < n {
		for j = i + 1; j < n && s[j] == s[i]; j++ {
		}
		cnt[s[i]-'0'], i = max(cnt[s[i]-'0'], j-i), j
	}
	return cnt[1] > cnt[0]
}
