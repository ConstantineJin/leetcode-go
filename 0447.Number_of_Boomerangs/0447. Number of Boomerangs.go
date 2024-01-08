package main

//func square(x int) int { return x * x } // 平方
//
//func factorial(x int) int { // 阶乘
//	if x == 0 {
//		return 1
//	}
//	var result = 1
//	for i := 1; i <= x; i++ {
//		result *= i
//	}
//	return result
//}
//
//func combinations(m, n int) int { // 组合数
//	if n > m {
//		return 0
//	}
//	return factorial(m) / (factorial(n) * factorial(m-n))
//}
//
//func numberOfBoomerangs(points [][]int) (ans int) {
//	var n = len(points)
//	var distances = make([]map[int]int, n)
//	for i := range distances {
//		distances[i] = make(map[int]int)
//	}
//	for i := range distances {
//		for j := i + 1; j < n; j++ {
//			var d = square(points[i][0]-points[j][0]) + square(points[i][1]-points[j][1])
//			distances[i][d]++
//			distances[j][d]++
//		}
//	}
//	for _, distance := range distances {
//		for _, d := range distance {
//			ans += combinations(d, 2)
//		}
//	}
//	return ans * 2
//}

func numberOfBoomerangs(points [][]int) (ans int) {
	var cnt = make(map[int]int) // 从同一个点出发，key：欧氏距离，value：数量
	for _, p1 := range points {
		clear(cnt) // 清空cnt
		for _, p2 := range points {
			var d2 = (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1]) // 双重循环求任意两点间的欧氏距离的平方和
			ans += cnt[d2]
			cnt[d2]++
		}
	}
	return ans * 2 // 题目规定点顺序不同算不同解（如i,j,k和i,k,j），故还需乘以2
}
