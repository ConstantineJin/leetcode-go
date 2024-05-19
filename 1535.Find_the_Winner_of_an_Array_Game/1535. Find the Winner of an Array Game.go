package main

func getWinner(arr []int, k int) int {
	mx, win := arr[0], 0
	for i := 1; i < len(arr) && win < k; i++ {
		if arr[i] > mx {
			mx = arr[i]
			win = 0
		}
		win++
	}
	return mx
}
