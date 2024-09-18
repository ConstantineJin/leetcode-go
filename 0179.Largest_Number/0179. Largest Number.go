package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) (ans string) {
	sort.Slice(nums, func(i, j int) bool { // 按照两数拼接结果的字典序降序排序
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	for _, num := range nums {
		ans += strconv.Itoa(num)
	}
	return ans
}
