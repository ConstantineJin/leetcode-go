package main

import (
	"math"
	"sort"
)

func minimumTime(time []int, totalTrips int) int64 {
	return int64(sort.Search(math.MaxInt, func(i int) bool {
		var cnt int
		for _, t := range time {
			cnt += i / t
			if cnt >= totalTrips {
				return true
			}
		}
		return false
	}))
}
