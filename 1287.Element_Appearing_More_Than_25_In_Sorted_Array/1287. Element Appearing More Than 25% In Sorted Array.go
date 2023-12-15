package main

func findSpecialInteger(arr []int) int {
	var n, l = len(arr), 1
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			l++
		} else {
			l = 1
		}
		if l > n/4 {
			return arr[i]
		}
	}
	return arr[0]
}
