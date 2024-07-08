package main

func numberOfAlternatingGroups(colors []int, k int) (ans int) {
	n := len(colors)
	colors = append(colors, colors[:k]...)
	length := make([]int, len(colors))
	for i, color := range colors {
		if i == 0 {
			continue
		}
		if color == colors[i-1] {
			length[i] = 1
		} else {
			length[i] = length[i-1] + 1
		}
	}
	for i := 0; i < k; i++ {
		length[i] = length[n+i]
	}
	for _, l := range length[:n] {
		if l >= k {
			ans++
		}
	}
	return
}
