package main

// 直接模拟
//func countPoints(rings string) (cnt int) {
//	rods := make([]map[string]bool, 10)
//	for i := range rods {
//		rods[i] = make(map[string]bool)
//	}
//	for i := 0; i < len(rings); i += 2 {
//		ring, _ := strconv.Atoi(string(rings[i+1]))
//		switch rings[i] {
//		case 'R':
//			rods[ring]["R"] = true
//		case 'G':
//			rods[ring]["G"] = true
//		case 'B':
//			rods[ring]["B"] = true
//		}
//	}
//	for _, rod := range rods {
//		if rod["R"] && rod["G"] && rod["B"] {
//			cnt++
//		}
//	}
//	return
//}

// 位运算
func countPoints(rings string) (cnt int) {
	rods := make([]int, 10)
	for i := 0; i < len(rings); i += 2 {
		rod := rings[i+1] - '0'
		switch rings[i] {
		case 'R':
			rods[rod] |= 1
		case 'G':
			rods[rod] |= 2
		case 'B':
			rods[rod] |= 4
		}
	}
	for _, rod := range rods {
		if rod == 7 {
			cnt++
		}
	}
	return
}
