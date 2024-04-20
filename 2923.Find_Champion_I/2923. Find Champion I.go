package main

//func findChampion(grid [][]int) (champion int) {
//	for i, j := 0, 0; i < len(grid); i++ {
//		// 如果某一列全为0就代表没有人击败过该队，则该列为冠军
//		for j = 0; j < len(grid[i]); j++ {
//			if grid[j][i] == 1 {
//				break
//			}
//		}
//		if j == len(grid[i]) {
//			return i
//		}
//	}
//	return
//}

func findChampion(grid [][]int) (ans int) {
	for i, row := range grid {
		if row[ans] == 1 {
			ans = i
		}
	}
	return
}
