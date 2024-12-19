package main

func maxChunksToSorted(arr []int) (ans int) {
	var mx int
	for i, x := range arr {
		mx = max(mx, x)
		if mx == i {
			ans++
		}
	}
	return
}
