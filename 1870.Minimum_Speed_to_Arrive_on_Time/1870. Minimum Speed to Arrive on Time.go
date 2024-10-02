package main

import (
	"math"
	"sort"
)

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if n > int(math.Ceil(hour)) {
		return -1
	}
	return sort.Search(1e7+1, func(speed int) bool {
		speed++ // 速度不能为 0
		var time float64
		for i, d := range dist {
			t := float64(d) / float64(speed)
			if i == n-1 {
				time += t
			} else {
				time += math.Ceil(t)
			}
		}
		return time <= hour
	}) + 1
}
