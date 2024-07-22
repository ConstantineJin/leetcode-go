package main

func minChanges(n int, k int) (ans int) {
	if k > n {
		return -1
	}
	for n > 0 || k > 0 {
		if (n&1) == 1 && (k&1) == 0 {
			ans++
		}
		if (n&1) == 0 && (k&1) == 1 {
			return -1
		}
		n >>= 1
		k >>= 1
	}
	return
}
