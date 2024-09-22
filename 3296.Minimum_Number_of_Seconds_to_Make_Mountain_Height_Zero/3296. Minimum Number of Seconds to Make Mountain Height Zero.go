package main

import (
	"math"
	"sort"
)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	return int64(sort.Search(math.MaxInt, func(timeLimit int) bool {
		var totalHeight int
		for _, time := range workerTimes {
			low, high := 0, mountainHeight
			for low <= high {
				mid := (low + high) / 2
				if timeLimit >= time*mid*(mid+1)/2 {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			totalHeight += high
			if totalHeight >= mountainHeight {
				return true
			}
		}
		return totalHeight >= mountainHeight
	}))
}
