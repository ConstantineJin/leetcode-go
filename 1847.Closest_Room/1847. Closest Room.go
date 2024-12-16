package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
	"slices"
)

func closestRoom(rooms [][]int, queries [][]int) []int {
	slices.SortFunc(rooms, func(a, b []int) int { return b[1] - a[1] }) // 将 rooms 按照 size 降序排序
	q := len(queries)
	queryIds := make([]int, q)
	for i := range queryIds {
		queryIds[i] = i
	}
	slices.SortFunc(queryIds, func(i, j int) int { return queries[j][1] - queries[i][1] }) // 将 queryIds 按照 minSize 降序排序
	ans := make([]int, q)
	for i := range ans {
		ans[i] = -1
	}
	roomIds := redblacktree.New[int, struct{}]()
	var j int
	for _, i := range queryIds {
		preferredId, minSize := queries[i][0], queries[i][1]
		for j < len(rooms) && rooms[j][1] >= minSize {
			roomIds.Put(rooms[j][0], struct{}{})
			j++
		}
		diff := math.MaxInt
		if node, ok := roomIds.Floor(preferredId); ok {
			diff = preferredId - node.Key
			ans[i] = node.Key
		}
		if node, ok := roomIds.Ceiling(preferredId); ok && node.Key-preferredId < diff {
			ans[i] = node.Key
		}
	}
	return ans
}
