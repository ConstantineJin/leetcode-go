package main

import (
	"slices"
	"strconv"
	"strings"
)

// isSimilar 判断两个组成数字频数完全相同的整数字符串是否“近似相等”
func isSimilar(x, y string) bool {
	if x == y { // 如果两数相等，直接返回 true
		return true
	}
	var diff int // 两数在同一位上数字不同的个数
	a, b := make([]byte, 4), make([]byte, 4)
	for i := range x {
		if x[i] != y[i] {
			if diff == 4 {
				return false
			}
			a[diff], b[diff] = x[i], y[i]
			diff++
		}
		if diff > 4 {
			return false
		}
	}
	return diff < 4 || a[0] == b[1] && b[0] == a[1] && a[2] == b[3] && b[2] == a[3] || a[0] == b[2] && b[0] == a[2] && a[1] == b[3] && b[1] == a[3] || a[0] == b[3] && b[0] == a[3] && a[1] == b[2] && b[1] == a[2] // 如果不同小于 4 处，必然可在两次交换内使两数相等；如果不同恰为 4 处，则当且仅当分成两组，每组组内交换后能使两数相等时两数“近似相等”
}

func countPairs(nums []int) (ans int) {
	n := len(nums)
	m := len(strconv.Itoa(slices.Max(nums))) // nums 中最大长度
	arr := make([]string, n)
	mp := make(map[string][]int)
	for i, num := range nums {
		a := strconv.Itoa(num)
		arr[i] = strings.Repeat("0", m-len(a)) + a // 在开头补 0 使数组中所有元素等长
		s := []byte(arr[i])
		slices.Sort(s)
		mp[string(s)] = append(mp[string(s)], i) // 将数组元素分成若干集合，每个集合中的数由频数相同的数字组成
	}
	for _, indices := range mp {
		for i := 0; i < len(indices); i++ {
			for j := i + 1; j < len(indices); j++ {
				if isSimilar(arr[indices[i]], arr[indices[j]]) {
					ans++
				}
			}
		}
	}
	return
}
