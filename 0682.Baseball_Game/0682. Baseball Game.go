package main

import "strconv"

func calPoints(operations []string) (ans int) {
	var points []int
	for _, operation := range operations {
		switch operation {
		case "+":
			points = append(points, points[len(points)-1]+points[len(points)-2])
		case "D":
			points = append(points, points[len(points)-1]*2)
		case "C":
			points = points[:len(points)-1]
		default:
			point, _ := strconv.Atoi(operation)
			points = append(points, point)
		}
	}
	for _, point := range points {
		ans += point
	}
	return
}
