package main

func findChampion(n int, edges [][]int) (champ int) {
	arr, find := make([]bool, n), false
	for _, edge := range edges {
		arr[edge[1]] = true
	}
	for i, v := range arr {
		if !v {
			if find {
				return -1
			}
			champ, find = i, true
		}
	}
	return
}
