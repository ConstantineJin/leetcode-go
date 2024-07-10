package main

import "math"

const inf = math.MaxInt

func minimumDistance(points [][]int) int {
	maxX1, maxX2, maxY1, maxY2 := -inf, -inf, -inf, -inf
	minX1, minX2, minY1, minY2 := inf, inf, inf, inf
	for _, point := range points {
		// 将曼哈顿距离转化为切比雪夫距离
		// 将每个点绕原点顺时针旋转45°，并扩大√2倍，则 (x, y) 变成 (x+y, y-x)
		// 将曼哈顿距离投影到x轴和y轴（选择较长的投影），会缩小√2倍，长度即为原来的曼哈顿距离
		x, y := point[0]+point[1], point[1]-point[0]

		if x > maxX1 { // x 最大次大
			maxX2 = maxX1
			maxX1 = x
		} else if x > maxX2 {
			maxX2 = x
		}
		if x < minX1 { // x 最小次小
			minX2 = minX1
			minX1 = x
		} else if x < minX2 {
			minX2 = x
		}
		if y > maxY1 { // y 最大次大
			maxY2 = maxY1
			maxY1 = y
		} else if y > maxY2 {
			maxY2 = y
		}
		if y < minY1 { // y 最小次小
			minY2 = minY1
			minY1 = y
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := inf
	for _, point := range points {
		x, y := point[0]+point[1], point[1]-point[0]
		dx, dy := f(x, maxX1, maxX2)-f(x, minX1, minX2), f(y, maxY1, maxY2)-f(y, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(v, v1, v2 int) int {
	if v == v1 {
		return v2
	}
	return v1
}
