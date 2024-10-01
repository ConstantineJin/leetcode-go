package main

func canArrange(arr []int, k int) bool {
	groups := make([]int, k)
	for _, v := range arr {
		groups[(v%k+k)%k]++
	}
	if groups[0]&1 != 0 {
		return false
	}
	for i := 1; i < (k+1)/2; i++ {
		if groups[i] != groups[k-i] {
			return false
		}
	}
	return true
}
