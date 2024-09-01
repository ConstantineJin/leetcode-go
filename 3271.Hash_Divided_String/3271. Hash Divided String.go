package main

func stringHash(s string, k int) (ans string) {
	for i := 0; i < len(s); i += k {
		var sum int
		for j := i; j < i+k; j++ {
			sum += int(s[j] - 'a')
		}
		ans += string('a' + byte(sum%26))
	}
	return
}
