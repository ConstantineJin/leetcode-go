package main

func numberOfAlternatingGroups(colors []int) (ans int) {
	colors = append(colors, colors[0], colors[1])
	for i := 1; i < len(colors)-1; i++ {
		if colors[i] != colors[i-1] && colors[i] != colors[i+1] {
			ans++
		}
	}
	return
}
