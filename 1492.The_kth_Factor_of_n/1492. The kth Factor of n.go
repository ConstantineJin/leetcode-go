package main

func kthFactor(n int, k int) int {
	var s, t []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			s, t = append(s, i), append([]int{n / i}, t...)
		}
	}
	if s[len(s)-1] == t[0] {
		t = t[1:]
	}
	if k <= len(s) {
		return s[k-1]
	} else if k > len(s)+len(t) {
		return -1
	} else {
		return t[k-len(s)-1]
	}
}
