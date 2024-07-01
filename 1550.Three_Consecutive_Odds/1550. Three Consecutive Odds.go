package main

func threeConsecutiveOdds(arr []int) bool {
	for i := 2; i < len(arr); i++ {
		if arr[i]%2 == 1 && arr[i-1]%2 == 1 && arr[i-2]%2 == 1 {
			return true
		}
	}
	return false
}
