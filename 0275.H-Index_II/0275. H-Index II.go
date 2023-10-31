package main

import "sort"

func hIndex(citations []int) int {
	return len(citations) - sort.Search(len(citations), func(i int) bool { return citations[i] >= len(citations)-i })
}
