package main

func minAnagramLength(s string) int {
	n := len(s)
next:
	for k := 1; k <= n/2; k++ {
		if n%k == 0 {
			var cnt0 [26]int
			for _, c := range s[:k] {
				cnt0[c-'a']++
			}
			for i := k * 2; i <= n; i += k {
				var cnt [26]int
				for _, c := range s[i-k : i] {
					cnt[c-'a']++
				}
				if cnt != cnt0 {
					continue next
				}
			}
			return k
		}
	}
	return n
}
