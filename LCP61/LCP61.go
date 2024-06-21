package main

import "cmp"

// 方法1:直接比较
//func temperatureTrend(temperatureA, temperatureB []int) (ans int) {
//	var cnt int
//	for i := 1; i < len(temperatureA); i++ {
//		if temperatureA[i] < temperatureA[i-1] && temperatureB[i] < temperatureB[i-1] || temperatureA[i] == temperatureA[i-1] && temperatureB[i] == temperatureB[i-1] || temperatureA[i] > temperatureA[i-1] && temperatureB[i] > temperatureB[i-1] {
//			cnt++
//			ans = max(ans, cnt)
//		} else {
//			cnt = 0
//		}
//	}
//	return
//}

// 方法2:库函数
func temperatureTrend(temperatureA, temperatureB []int) (ans int) {
	var cnt int
	for i := 1; i < len(temperatureA); i++ {
		if cmp.Compare(temperatureA[i], temperatureA[i-1]) == cmp.Compare(temperatureB[i], temperatureB[i-1]) {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 0
		}
	}
	return
}
